#!/bin/bash

echo "Locating the working directory for this script"
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="${DIR}/${SOURCE}"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

echo "Script directory: $DIR"

echo "Available artifacts:"
ls -al $DIR/artifacts/

BUILD_VERSION_NUMBER=$(head -n 1 $DIR/artifacts/.semver)
echo "BUILD_VERSION_NUMBER = $BUILD_VERSION_NUMBER"

# Suffix passed in by the caller
DEPLOY_ENV_SUFFIX=$1
echo "DEPLOY_ENV_SUFFIX = $DEPLOY_ENV_SUFFIX"

# Added some complexity to allow for running locally
DOCKER_IMAGE_REPO=${TRAVIS_REPO_SLUG}
if [ -z "$DOCKER_IMAGE_REPO"]; then
  echo "Using local git context for DOCKER_IMAGE_REPO"
  DOCKER_IMAGE_REPO=`git remote -v | head -n1 | awk '{print $2}' | sed "s/.*://" | sed "s/\.git//"`
fi

if [ ! -z "$DEPLOY_ENV_SUFFIX" ]; then
  # We will push to a different docker repo
  echo "Adding suffix $DEPLOY_ENV_SUFFIX to $DOCKER_IMAGE_REPO"
  DOCKER_IMAGE_REPO="${DOCKER_IMAGE_REPO}-${DEPLOY_ENV_SUFFIX}"
fi

DOCKER_IMAGE_NAME="${DOCKER_IMAGE_REPO}:${BUILD_VERSION_NUMBER}"

echo "Building docker image $DOCKER_IMAGE_NAME"
docker build -t $DOCKER_IMAGE_NAME $DIR

# Only push if the env variables are set
if [ ! -z $DOCKER_USER ] && [ ! -z $DOCKER_PASSWORD ]; then
  echo "Logging in to dockerhub"
  docker login -u="$DOCKER_USER" -p="$DOCKER_PASSWORD"
  echo "Pushing $DOCKER_IMAGE_NAME"
  docker push $DOCKER_IMAGE_NAME

  DOCKER_IMAGE_LATEST="$DOCKER_IMAGE_REPO:latest"
  echo "Tagging $DOCKER_IMAGE_NAME as latest"
  docker tag $DOCKER_IMAGE_NAME $DOCKER_IMAGE_LATEST
  echo "Pushing $DOCKER_IMAGE_LATEST"
  docker push $DOCKER_IMAGE_LATEST

  docker logout
else
  if [ ! -z $TRAVIS ]; then
    echo "Deployment Error: docker login credentials not present"
    exit 1
  fi
fi

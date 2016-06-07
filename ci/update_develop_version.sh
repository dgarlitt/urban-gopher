#!/bin/bash

if [ $TRAVIS_BRANCH = "master" ] && [ $TRAVIS_PULL_REQUEST = "false" ]; then
  echo "Adding new remote for pushing to GitHub"
  git remote add github "git@github.com:${TRAVIS_REPO_SLUG}.git"

  echo "Checking out the develop branch"
  git checkout -b develop

  echo "Updating the version file"
  cat .semver | awk 'BEGIN { FS="."; OFS="."; } { print $1,$2,++$3 > ".semver" }'

  echo "Adding version file to index"
  git add .semver

  echo "Committing new version number: $(head -n1 .semver)"
  git commit -m "Auto-increment version number"

  echo "Pushing version change commit to remote"
  git push github develop
else
  echo "No post-deploy action required"
fi

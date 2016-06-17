[![Build Status](https://travis-ci.org/dgarlitt/urban-gopher.svg)](https://travis-ci.org/dgarlitt/urban-gopher)
[![Coverage Status](https://coveralls.io/repos/github/dgarlitt/urban-gopher/badge.svg?branch=master)](https://coveralls.io/github/dgarlitt/urban-gopher?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dgarlitt/urban-gopher)](https://goreportcard.com/report/github.com/dgarlitt/urban-gopher)

![Urban Gopher Logo](https://raw.githubusercontent.com/dgarlitt/image-repo/master/urban-gopher/urban-gopher-art.jpg)

# Urban Dictionary Definition Lookup Service

This is a simple service that will look-up the "best" definition for a given term
in Urban Dictionary using the Mashape API. You will need to provide a Mashape
API Key to make a request from this service.

A Mashape API key can be provided as an environment variable:

```
export URBAN_GOPHER_API_KEY=<your-mashape-api-key>
```

Then you can startup the service and hit the endpoint as follows:

```
curl http://localhost:8008/define?term=wat
```

Alternatively, you can provide the API key as a header when making a request:

Example request:

```
curl -H "X-API-Key: <your-mashape-api-key>" http://localhost:8008/define?term=wat
```

Mashape API keys will not be logged in the log output.

## Install and Run

```sh
go get github.com/dgarlitt/urban-gopher
cd $GOPATH/src/github.com/dgarlitt/urban-gopher
go build -v -race -o ci/artifacts/urban-gopher
ci/artifacts/urban-gopher
```

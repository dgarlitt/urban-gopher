[![Build Status](https://travis-ci.org/dgarlitt/urban-gopher.svg)](https://travis-ci.org/dgarlitt/urban-gopher)
[![Coverage Status](https://coveralls.io/repos/github/dgarlitt/urban-gopher/badge.svg?branch=master)](https://coveralls.io/github/dgarlitt/urban-gopher?branch=master)

# Urban Dictionary Definition Lookup Service

This is a simple service that will look-up the "best" definition for a given term
in Urban Dictionary using the Mashape API. You will need to provide a Mashape
API Key to make a request from this service.

Mashape API keys will not be logged in the log output.

Example request:

```
curl -H "X-API-Key: <your-mashape-api-key>" http://localhost:8008/define?term=wat
```

## Install and Run

```sh
go get github.com/dgarlitt/urban-gopher
cd $GOPATH/src/github.com/dgarlitt/urban-gopher
go build -v -race -o deploy/artifacts/urban-gopher
deploy/artifacts/urban-gopher
```

![Golang Gopher Logo](https://raw.githubusercontent.com/dgarlitt/image-repo/master/tech-logos/golang-gopher.png)

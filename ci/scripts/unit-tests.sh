#!/bin/bash

set -e -x
 
export GOPATH=$PWD
mkdir -p src/github.com/JeffDeCola/
cp -R ./RESTful-API-test src/github.com/JeffDeCola/.

go test -v -cover github.com/JeffDeCola/RESTful-API-test/...
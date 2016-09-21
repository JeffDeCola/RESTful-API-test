#!/bin/bash
# RESTful-API-test unit-tset.sh

set -e -x

# The code is located in /RESTful-API-test
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./RESTful-API-test to $GOPATH/src/github.com/JeffDeCola/RESTful-API-test
mkdir -p src/github.com/JeffDeCola/
cp -R ./RESTful-API-test src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/JeffDeCola/RESTful-API-test

# RUN unit_tests
# go test -v -cover ./...
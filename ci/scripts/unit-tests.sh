#!/bin/bash

set -e -x

go get ./...
go test -v -cover ./RESTful-API-test
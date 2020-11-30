# RESTful-API-test

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/RESTful-API-test)](https://goreportcard.com/report/github.com/JeffDeCola/RESTful-API-test)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/RESTful-API-test?status.svg)](https://godoc.org/github.com/JeffDeCola/RESTful-API-test)
[![Maintainability](https://api.codeclimate.com/v1/badges/57a79ec6bf13f735c4bf/maintainability)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A very limited RESTful API in which you can GET
and POST data from a database via a CLI http client._

[GitHub Webpage](https://jeffdecola.github.io/RESTful-API-test/)
_built with
[concourse ci](https://github.com/JeffDeCola/RESTful-API-test/blob/master/ci-README.md)_

## ORIGINAL SOURCE

Original go code from [phalt](https://github.com/phalt).

Written in go using gin &amp; gorp.

## RUN

From Command line

```bash
go run main.go
```

In another terminal, use a CLI http client like httpie and you can do the
following commands:

### Create New Article

```bash
http POST localhost:8000/articles title="A simple RESTful-API-test" content="Hello-World"
```

### Query Entry List - Returns all articles in the list

```bash
http localhost:8000/articles
```

### Query Single Resource - Get back single article via it's id

```bash
http localhost:8000/articles/1
```

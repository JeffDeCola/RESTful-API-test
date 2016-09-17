# RESTful-API-test

[![Code Climate](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/badges/gpa.svg)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)
[![Go Report Card](https://goreportcard.com/badge/jeffdecola/RESTful-API-test)](https://goreportcard.com/report/jeffdecola/RESTful-API-test)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/RESTful-API-test?status.svg)](https://godoc.org/github.com/JeffDeCola/RESTful-API-test)


`RESTful-API-test` is a very limited RESTful API.

Written in go using gin &amp; gorp.

It is unit tested and built using concourse ci.

(Original Code from https://github.com/phalt)

----

## USAGE

go run main.go

In another terminal, use a cli http client like httpie.

### CHECK BODY

```bash
http localhost:8000/ --body
```

### CREATE NEW ARTICLE

```bash
http POST localhost:8000/articles title="A simple API in Go" content="This is my content"
```

### QUERY ENTIRE LIST - Returns all articles in the list

```bash
http localhost:8000/articles
```

### QUERY SINGLE RESOURCE - Get back single article via it's id

```bash
http localhost:8000/articles/1
```

## CONCOURSE CI

A concourse ci pipeline is used to unit test.

![IMAGE - hello-go concourse ci piepline - IMAGE](docs/RESTful-API-test-pipeline.jpg)

A /ci/.credentials file needs to be created for your slack_url and repo_github_token.

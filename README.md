# RESTful-API-test

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/RESTful-API-test)](https://goreportcard.com/report/github.com/JeffDeCola/RESTful-API-test)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/RESTful-API-test?status.svg)](https://godoc.org/github.com/JeffDeCola/RESTful-API-test)
[![Maintainability](https://api.codeclimate.com/v1/badges/57a79ec6bf13f735c4bf/maintainability)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/RESTful-API-test/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org

`RESTful-API-test` is a very limited RESTful API in which you can GET
and POST data from a database via a CLI http client.

[RESTful-API-test GitHub Webpage](https://jeffdecola.github.io/RESTful-API-test/)

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

## UPDATE GITHUB WEBPAGE & UNIT TESTS USING CONCOURSE (OPTIONAL)

For fun, I use concourse to  update
[RESTful-API-test GitHub Webpage](https://jeffdecola.github.io/RESTful-API-test/),
run my unit-tests and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - RESTful-API-test concourse ci pipeline - IMAGE](docs/pics/RESTful-API-test-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/scripts/readme-github-pages.sh).
* `job-unit-tests` runs task
  [unit-tests.sh](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/scripts/unit-tests.sh).

The concourse `resources types` are,

* `RESTful-API-test` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).

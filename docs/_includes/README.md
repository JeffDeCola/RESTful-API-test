
# ORIGINAL SOURCE

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

## UNIT TESTING AND MY GITHUB WEBPAGE IS UPDATED USING CONCOURSE

For fun, I use concourse to automate unit testing, update
[RESTful-API-test GitHub Webpage](https://jeffdecola.github.io/RESTful-API-test/)
and alert me of the changes via repo status and slack.

The unit testing is accomplished by running this script this script
[here](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/scripts/unit-tests.sh).

The github webpage update is accomplished this by copying and editing
this `README.md` file to `/docs/_includes/README.md`.
You can see the concourse task (a shell script)
[here](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/scripts/readme-github-pages.sh).

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/RESTful-API-test/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - RESTful-API-test concourse ci pipeline - IMAGE](pics/RESTful-API-test-pipeline.jpg)

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).

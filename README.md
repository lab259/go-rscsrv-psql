# go-rscsrv-psql [![CircleCI](https://circleci.com/gh/lab259/go-rscsrv-psql.svg?style=shield)](https://circleci.com/gh/lab259/go-rscsrv-psql) [![Go Report Card](https://goreportcard.com/badge/github.com/lab259/go-rscsrv-psql)](https://goreportcard.com/report/github.com/lab259/go-rscsrv-psql) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=shield)](http://godoc.org/github.com/lab259/go-rscsrv-psql) [![Release](https://img.shields.io/github/release/lab259/go-rscsrv-psql.svg?style=shield)](https://github.com/lab259/go-rscsrv-psql/releases/latest)

## Getting Started

### Prerequisites

What things you need to setup the project:

- [go](https://golang.org/doc/install)
- [ginkgo](http://onsi.github.io/ginkgo/)

### Environment

For start developing you must clone the project:

```bash
git clone git@github.com:lab259/go-rscsrv-psql.git
```

Now, the dependencies must be installed.

```
go mod download
```

:wink: Finally, you are done to start developing.

### Running tests

In the root directory, start PostgreSQL (using [`docker-compose`](https://docs.docker.com/compose/install/)):

```bash
docker-compose up -d
```

Then execute:

```bash
make test
```

To enable coverage, execute:

```bash
make coverage
```

To generate the HTML coverage report, execute:

```bash
make coverage coverage-html
```

# go-rscsrv-psql [![CircleCI](https://circleci.com/gh/lab259/go-rscsrv-psql.svg?style=shield)](https://circleci.com/gh/lab259/go-rscsrv-psql) [![Go Report Card](https://goreportcard.com/badge/github.com/lab259/go-rscsrv-psql)](https://goreportcard.com/report/github.com/lab259/go-rscsrv-psql) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=shield)](http://godoc.org/github.com/lab259/go-rscsrv-psql) [![Release](https://img.shields.io/github/release/lab259/go-rscsrv-psql.svg?style=shield)](https://github.com/lab259/go-rscsrv-psql/releases/latest)

## Getting Started

### Prerequisites

What things you need to setup the project:

- [go](https://golang.org/doc/install)
- [golang/dep](https://github.com/golang/dep)
- [ginkgo](http://onsi.github.io/ginkgo/)

### Environment

For start developing you must create a `GOPATH` structure:

```
+-- /
|---- src
|------ github.com
|-------- lab259
|---------- go-rscsrv-psql <- Here is where you will clone this repository.
```

Use the following command:

```bash
mkdir -p src/github.com/lab259/go-rscsrv-psql && git clone git@github.com:lab259/go-rscsrv-psql.git src/github.com/lab259/go-rscsrv-psql
```

Now, the dependencies must be installed.

```
cd src/github.com/lab259/go-rscsrv-psql && make dep-ensure
```

:wink: Finally, you are done to start developing.

### Running tests

In the `src/github.com/lab259/go-rscsrv-psql` directory, execute:

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

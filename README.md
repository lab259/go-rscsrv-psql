# go-package-boilerplate [![CircleCI](https://circleci.com/gh/lab259/go-package-boilerplate.svg?style=shield&circle-token=224f68e222b4a6abeb01f2d0dda3b4cf264b806e)](https://circleci.com/gh/lab259/go-package-boilerplate)

> See here [how to create a repository from a template](https://help.github.com/en/articles/creating-a-repository-from-a-template)

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
|---------- go-package-boilerplate <- Here is where you will clone this repository.
```

Use the following command:

```bash
mkdir -p src/github.com/lab259/go-package-boilerplate && git clone git@github.com:lab259/go-package-boilerplate.git src/github.com/lab259/go-package-boilerplate
```

Now, the dependencies must be installed.

```
cd src/github.com/lab259/go-package-boilerplate && make dep-ensure
```

:wink: Finally, you are done to start developing.

### Running tests

In the `src/github.com/lab259/go-package-boilerplate` directory, execute:

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

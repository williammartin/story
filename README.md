## Story

Story is a Go implementation of the Storyscript Cloud CLI

### Getting started

To get story:

`git clone git@github.com:williammartin/story.git`

To build:

`make build`

### Running the tests

Install the `ginkgo` test runner:

`go get -u github.com/onsi/ginkgo/ginkgo`

To run the acceptance tests you will currently need to set an env var `STORYSCRIPT_TOKEN`. Additionally, the acceptance tests assume a fresh environment.

The test targets are:

```
make test
make test-units
make test-acceptance
```

If you really don't want to install ginkgo, you can use the default go test runner:

```
go test ./...
```

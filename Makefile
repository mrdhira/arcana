VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOFILES := $(wildcard *.go)
STIME := $(shell date +%s)

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## serveHttp: start http server without docker
serve-http:
	@echo " > Starting Serve Http Program..."
	go run main.go serveHttp
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"
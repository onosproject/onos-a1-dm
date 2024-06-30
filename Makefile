# SPDX-License-Identifier: Apache-2.0
# Copyright 2019 Open Networking Foundation
# Copyright 2024 Intel Corporation

.PHONY: build license

GOLANG_CI_VERSION := v1.52.2

all: build

build: # @HELP build the Go library
build: schemas go-build

test: # @HELP run the unit tests and source code validation
test: build lint license

go-build:
	cd go && go build ./... && gofmt -s -w .

schemas:
	cd ./build/bin/ && ./compile-a1-schemas.sh

lint: # @HELP examines Go source code and reports coding problems
	cd ./go; \
	golangci-lint --version | grep $(GOLANG_CI_VERSION) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b `go env GOPATH`/bin $(GOLANG_CI_VERSION); \
	golangci-lint run --timeout 15m

license: # @HELP run license checks
	rm -rf venv
	python3 -m venv venv
	. ./venv/bin/activate;\
	python3 -m pip install --upgrade pip;\
	python3 -m pip install reuse;\
	reuse lint

check-version: # @HELP check version is duplicated
	./build/bin/version_check.sh all

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '

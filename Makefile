# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0


.PHONY: build

all: schemas golang

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

test: # @HELP run the unit tests and source code validation
test: golang schemas

jenkins-test: # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: jenkins-tools

golang:
	cd go && go build ./...

schemas:
	cd ./build/bin/ && ./compile-a1-schemas.sh

publish: # @HELP publish version on github
	./build/build-tools/publish-version ${VERSION}
	./build/build-tools/publish-version go/${VERSION}

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

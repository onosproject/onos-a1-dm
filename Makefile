# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build

all: schemas golang

golang:
	cd go && go build ./...

schemas:
	cd ./build/bin/ && ./compile-a1-schemas.sh
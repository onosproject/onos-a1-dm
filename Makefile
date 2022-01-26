.PHONY: build

all: schemas golang

golang:
	cd go && go build ./...

schemas:
	cd ./build/bin/ && ./compile-a1-schemas.sh
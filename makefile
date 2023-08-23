## Example makefile for Go with fuzz test

## This is for not using go.mod
# .EXPORT_ALL_VARIABLES:
# GOPATH=$(shell cd)
# GO111MODULE=off

default: run

.PHONY: help
help:
	$(info Usage:)
	$(info make help  - this help)
	$(info make qc    - go vet )
	$(info make lint  - All linter and error checkors )
	$(info make run   - run with go vet)
	$(info make test  - run tests )
	$(info make bench - run bench)
	$(info make build - build with full lint and staticcheck)
	$(info make clean - delete *.exe)

qc: 
	@-go vet .
	@echo -----------------------------------------------------------------

lint: qc
	-revive -formatter stylish . 
	-errcheck . 
	@-echo -----------------------------------------------------------------

test: qc
	-go test ./...
	@echo -----------------------------------------------------------------

bench: qc 
	go test -benchmem -run=. -bench=. -benchtime=20s
	@echo -----------------------------------------------------------------
	go test -fuzz=./... -fuzztime=20s 
	@echo -----------------------------------------------------------------
	
build: lint test 
	go build -gcflags="-m=2" --ldflags="-s -w -race" -trimpath .
	@echo -----------------------------------------------------------------

run: qc
	go run .
	@echo -----------------------------------------------------------------

.PHONY : clean
clean :
	@go clean .
	@-rm *.exe      

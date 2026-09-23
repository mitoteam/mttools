# prepare variables
APP_VERSION := $(file < VERSION)

# Run all tests
.PHONY: tests
tests::
	clear
	go test

# target alias: test = tests
.PHONY: test
test: tests


.PHONY: version
version:
	@echo "mttools version from 'VERSION' file: '${APP_VERSION}'"


.PHONY: deps
deps:
	go get -u -t ./...
	go mod tidy

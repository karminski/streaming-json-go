.PHONY: test clean

test:
	go test -v

benchmark:
	go test -bench=.

test-cover:
	go test -cover --count=1

test-cover-report:
	PROJECT_PWD=$(shell pwd) go test -coverprofile cover.out 
	go tool cover -html=cover.out -o cover.html

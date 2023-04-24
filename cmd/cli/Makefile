
SHELL := /bin/bash

build:
	source env.sh && go build -o bin/app

run: build
	source env.sh && ./bin/app --repository-owner $(o) --repository-name $(n) --pull-request-number $(pr)

test:
	go test -v ./... -count=1

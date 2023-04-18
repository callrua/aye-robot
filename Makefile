
SHELL := /bin/bash

build:
	source env.sh && go build -o bin/app

run: build
	source env.sh && ./bin/app

test: build
	source env.sh && go test -v ./... -count=1

docker-build:
	source env.sh && docker build -t aye-robot:latest .

docker-run: docker-build
	source env.sh && docker run --rm -p 32768:9090 aye-robot

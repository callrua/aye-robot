
SHELL := /bin/bash

test:
	go test -v ./... -count=1

docker-build-server:
	source env.sh && docker build -t aye-robot:latest .

docker-run-server: docker-build
	source env.sh && docker run --rm -p 32768:9090 aye-robot

REVISION ?= $(shell git rev-parse --short HEAD)

container:
	docker build -t kazukousen/go-api-utils:$(REVISION) .

push:
	docker push kazukousen/go-api-utils

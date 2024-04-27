PROJECT=trouble-shooting
BINARY=trouble-shooting
REGISTRY=ccr.ccs.tencentyun.com
OUTPUT:=bin
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
ifdef ${ts_registry}
	echo "registry addr: ${ts_registry}"
	REGISTRY="ccr.ccs.tencentyun.com"
else
	REGISTRY=${ts_registry}
endif

base:
	mkdir -p ${OUTPUT}

GIT_TAG := $(shell git describe --tags --abbrev=0)

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	

default:
	go build -o ${OUTPUT}/${BINARY} main.go
linux-amd64:
	echo "ts version: ${GIT_TAG}"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${OUTPUT}/${BINARY} main.go
linux-arm:
	echo "ts version: ${GIT_TAG}"
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ${OUTPUT}/${BINARY} main.go
test:
	go test -v -coverprofile=cover.out ./router/
	go tool cover -html=cover.out -o test.html

docker:
	echo "ts image tag: ${REGISTRY}/${PROJECT}/ts:${GIT_TAG}"
	docker build -t ${REGISTRY}/${PROJECT}/:${GIT_TAG} .

podman: 
	echo "ts image tag: ${REGISTRY}/${PROJECT}/ts:${GIT_TAG}"
	podman build -t ${REGISTRY}/${PROJECT}/ts:${GIT_TAG} .
clean:
	rm -rf ${OUTPUT}

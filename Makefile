PROJECT=trouble-shooting
BINARY=trouble-shooting
REGISTRY=ccr.ccs.tencentyun.com
OUTPUT:=bin
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
TS_REGISTRY ?= ccr.ccs.tencentyun.com
# ifdef ${ts_registry}
# 	echo "registry addr: ${ts_registry}"
# 	REGISTRY="ccr.ccs.tencentyun.com"
# else
# 	REGISTRY=${ts_registry}
# endif

base:
	mkdir -p ${OUTPUT}

TAG := $(shell cat VERSION)

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	

default:
	go build -o ${OUTPUT}/${BINARY} main.go
linux-amd64:
	echo "ts version: ${TAG}"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${OUTPUT}/${BINARY} main.go
linux-arm:
	echo "ts version: ${TAG}"
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ${OUTPUT}/${BINARY} main.go
test:
	go test -v -coverprofile=cover.out ./router/
	go tool cover -html=cover.out -o test.html

docker:
	echo "ts image tag: ${TS_REGISTRY}/${PROJECT}/ts:${TAG}"
	docker build -t ${TS_REGISTRY}/${PROJECT}/:${TAG} .

podman: 
	echo "ts image tag: ${TS_REGISTRY}/${PROJECT}/ts:${TAG}"
	podman build -t ${REGISTRY}/${PROJECT}/ts:${TAG} .
clean:
	rm -rf ${OUTPUT}

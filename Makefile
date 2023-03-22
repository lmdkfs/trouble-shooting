PROJECT="trouble-shooting"
BINARY="trouble-shooting"
REGISTRY="ccr.ccs.tencentyun.com"
ifdef ${tsregistry}
	echo "registry addr: ${tsregistry}"
	REGISTRY="ccr.ccs.tencentyun.com"
else
	REGISTRY=${tsregistry}
endif

ifdef ${tsversion}
	echo "ts version: ${tsversion}"
        VERSION=0.2
else 
	VERSION=$(tsversion)
endif

default:
	go build -o ./bin/${BINARY} main.go
linux:
	echo "ts version: ${VERSION}"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/${BINARY} main.go
test:
	go test -v -coverprofile=cover.out ./router/
	go tool cover -html=cover.out -o test.html

docker:
	echo "ts version: ${VERSION}"
	docker build -t ${REGISTRY}/${PROJECT}/:${VERSION} .

podman: 
	echo "ts version: ${VERSION}"
	podman build -t ${REGISTRY}/${PROJECT}/ts:${VERSION} .
clean:
	cd bin && if [ -f ${BINARY} ] ; then rm  ${BINARY} ; fi

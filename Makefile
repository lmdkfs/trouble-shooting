PROJECT="trouble-shooting"
BINARY="trouble-shooting"
VERSION=0.2
default:
	go build -o ./bin/${BINARY} main.go
linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/${BINARY} main.go
test:
	go test -v -coverprofile=cover.out ./router/
	go tool cover -html=cover.out -o test.html

docker:
	docker build -t ${PROJECT}:${VERSION} .

podman: 
	podman build -t ${PROJECT}:${VERSION} .
clean:
	cd bin && if [ -f ${BINARY} ] ; then rm  ${BINARY} ; fi

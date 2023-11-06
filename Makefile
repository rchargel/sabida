BINARY_NAME=sabida

all: build

pre-build:
	cd src && sqlc generate
	
build-linux-amd64: pre-build
	cd src && GOARCH=amd64 GOOS=linux go build -o ../${BINARY_NAME}-linux-amd64 .
	
build-linux-arm64: pre-build
	cd src && GOARCH=arm64 GOOS=linux go build -o ../${BINARY_NAME}-linux-arm64 .
	
build-mac-amd64: pre-build
	cd src && GOARCH=amd64 GOOS=darwin go build -o ../${BINARY_NAME}-darwin-amd64 .
	
build-mac-arm64: pre-build
	cd src && GOARCH=arm64 GOOS=darwin go build -o ../${BINARY_NAME}-darwin-arm64 .
	
build-windows: pre-build
	cd src && GOARCH=amd64 GOOS=windows go build -o ../${BINARY_NAME}-windows-amd64 .

build: build-linux-amd64 build-linux-arm64 build-mac-amd64 build-mac-arm64 build-windows
	
run: build
	./${BINARY_NAME}
	
build-image: pre-build
	docker build --tag sabida .

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin-amd64
	rm -f ${BINARY_NAME}-linux-amd64
	rm -f ${BINARY_NAME}-windows-amd64
	rm -f ${BINARY_NAME}-darwin-arm64
	rm -f ${BINARY_NAME}-linux-arm64
	
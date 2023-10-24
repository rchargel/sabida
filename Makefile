BINARY_NAME=sabida

all: build

pre-build:
	sqlc generate

build: pre-build
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 .
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux-amd64 .
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows-amd64 .
	GOARCH=arm GOOS=darwin go build -o ${BINARY_NAME}-darwin-arm .
	GOARCH=arm GOOS=linux go build -o ${BINARY_NAME}-linux-arm .
	
run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin-amd64
	rm -f ${BINARY_NAME}-linux-amd64
	rm -f ${BINARY_NAME}-windows-amd64
	rm -f ${BINARY_NAME}-darwin-arm
	rm -f ${BINARY_NAME}-linux-arm
	
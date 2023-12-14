BINARY_NAME=echoserver
BIN_DIR=bin

all: clean build copyscript

build:

	GOARCH=amd64 GOOS=linux go build -o ./${BIN_DIR}/linux/${BINARY_NAME}  cmd/echoserver/main.go
	GOARCH=amd64 GOOS=windows go build -o ./${BIN_DIR}/win/${BINARY_NAME}.exe  cmd/echoserver/main.go

copyscript:
	cp ./script/run.bat ./${BIN_DIR}/win
	cp ./script/run.sh ./${BIN_DIR}/linux

clean:
	rm -rf ./${BIN_DIR}
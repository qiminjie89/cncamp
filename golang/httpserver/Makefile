export tag=v1.0
root:
	export ROOT=github.com/qiminjie89/cncamp/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH:=amd64 go build -o bin/amd64 .

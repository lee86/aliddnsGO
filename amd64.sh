#!/bin/bash
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -ldflags="-s -w" -a -v -o aliddns ./
scp -P1022 -O aliddns root@192.168.50.1:/root
go env -w GOOS=darwin
go env -w GOARCH=arm64

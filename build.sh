#!/bin/sh
set -v on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -gcflags "all=-N -l" -o ../order ./main.go
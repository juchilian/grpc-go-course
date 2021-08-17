#!/bin/sh

go mod download

find handlers -name main.go -type f \
 | xargs -n 1 dirname \
 | xargs -n 1 -I@ bash -c "cd ./@ && packr2 && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -installsuffix cgo -o main . && pwd && ls"
#!/usr/bin/env sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/drone-rocket
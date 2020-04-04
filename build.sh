#!/usr/bin/env bash
rm -rf kuake
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=kuake

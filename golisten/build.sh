#!/usr/bin/env zsh
go get -v all

GOOS=linux go build -o build/listen src/listen.go
#zip -jrm build/main.zip build/main

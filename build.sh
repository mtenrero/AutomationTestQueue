#!/bin/bash

# CONTROLLER
env GOOS=windows GOARCH=amd64 go build -o releases/atq-win-amd64.exe 
env GOOS=windows GOARCH=386 go build -o releases/atq-win-i386.exe 

env GOOS=linux GOARCH=386 go build -o releases/atq-linux-i386 
env GOOS=linux GOARCH=amd64 go build -o releases/atq-linux-amd64 

env GOOS=darwin GOARCH=amd64 go build -o releases/atq-osx-amd64 
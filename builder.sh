#!/bin/sh

# linux 64bit
GOOS=linux GOARCH=amd64 go build -o webhook_linux_64bit
upx -9 webhook_linux_64bit

# linux 32bit
GOOS=linux GOARCH=386 go build -o webhook_linux_32bit
upx -9 webhook_linux_32bit

# windows 64bit
GOOS=windows GOARCH=amd64 go build -o webhook_64bit.exe
upx -9 webhook_64bit.exe

# windows 32bit
GOOS=windows GOARCH=386 go build -o webhook_32bit.exe
upx -9 webhook_32bit.exe

# Mac OS X 64bit
GOOS=darwin GOARCH=amd64 go build -o webhook_mac
upx -9 webhook_mac

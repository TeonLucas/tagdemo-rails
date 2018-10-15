#!/bin/sh

# Go installed
if [ -z "$GOPATH" ]; then
   echo "Please set GOPATH"
   exit
fi

# Get dependencies
if ! [ -d "$GOPATH/src/github.com/DavidSantia/tag_api" ]; then
   go get "github.com/DavidSantia/tag_api"
fi
if ! [ -d "$GOPATH/src/github.com/DavidHuie/quartz/go/quartz" ]; then
   go get "github.com/DavidHuie/quartz/go/quartz"
fi

# Build and install goserver
cd goserver
go build
mv goserver ../bin

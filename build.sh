#!/bin/sh

export GOPATH=$(pwd)

#dependencies
go get labix.org/v2/mgo

#build
go install github.com/user/hello


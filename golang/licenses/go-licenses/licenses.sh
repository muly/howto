#!/bin/sh

set -xe

## install go licenses in the go bin
(
    if [ -z "$GOPATH"]; then
        export GOPATH=$(go env GOPATH)
    fi

    echo $GOPATH

    mkdir -p $GOPATH/src
    cd $GOPATH/src

    go get github.com/google/go-licenses
)

## run go licenses
which go-licenses
go-licenses csv github.com/muly/howto/golang/licenses/go-licenses > licenses.csv
sort -o licenses.csv licenses.csv

## if the license file is changes, report failure indicating that the file needs to be "git added"
if [[ $(git diff --name-only | grep licenses.csv) ]]; then
    git diff licenses.csv
    echo "license file needs update"
    exit 1
fi

echo "success"
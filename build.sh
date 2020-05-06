#!/usr/bin/env bash

echo "building...";
go build -o gitscan *.go;

echo "copying...";
cp gitscan $GOPATH/bin;
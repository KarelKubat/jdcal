#!/bin/sh

set -v

zero="main/makefullmoons/zero-fullmoons-go.txt"
out="fullmoons.out"
go="fullmoons.go"

cp "$zero" "$go" || exit 1
go run main/makefullmoons/makefullmoons.go "$out" || exit 1
mv "$out" "$go" || exit 1
gofmt -w "$go" || exit 1


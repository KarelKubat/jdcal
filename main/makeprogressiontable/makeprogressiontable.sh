#!/bin/sh

out="main/makeprogressiontable/progressiontable.out"
go="progressiontable.go"

go run main/makeprogressiontable/makeprogressiontable.go "$out" || exit 1
gofmt -w "$out" || exit 1
mv "$out" "$go" || exit 1

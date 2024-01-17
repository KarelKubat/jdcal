#!/bin/sh

tmpfile=main/bigconversiontest/testdates/testdates.out
gofile=main/bigconversiontest/testdates/testdates.go

if [ -f "$gofile" ] ; then
    echo "$gofile already exists, won't overwrite"
    exit 1
fi
rm -f "$tmpfile"

go run main/bigconversiontest/maketestdates/maketestdates.go "$tmpfile" || exit 1
mv "$tmpfile" "$gofile" || exit 1
gofmt -w "$gofile" || exit 1

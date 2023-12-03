package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	hdr = `package jdcal

// Generated file, don't edit by hand. Look for a nearby Makefile.

import (
	"time"
)

type Entry struct {
	JDate, GDate Date
}

// [...] is syntactic sugar to let the compiler figure out the array size. That way
// we get a fixed size array and not a slice.
var Table = [...]Entry{
`

	ftr = `
}`
)

var output bytes.Buffer

func main() {
	if len(os.Args) != 3 {
		check(errors.New("usage: go run what/ever/maketable.go INPUT OUTPUT"))
	}
	content, err := os.ReadFile(os.Args[1])
	check(err)

	splitRe := regexp.MustCompile(`[ \t]+`)

	out(hdr)

	lineno := 0
	for _, line := range strings.Split(string(content), "\n") {
		lineno++
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		parts := splitRe.Split(line, -1)

		yr := atoi(parts[0])
		jmonth := month(parts[1])
		jday := atoi(parts[2])
		gmonth := month(parts[3])
		gday := atoi(parts[4])

		out("{\n")
		out("JDate: Date{Year: %v, Month: time.%v, Day: %v, Type: Julian},\n", yr, jmonth, jday)
		out("GDate: Date{Year: %v, Month: time.%v, Day: %v, Type: Gregorian},\n", yr, gmonth, gday)
		if len(parts) == 6 {
			out("// Diff: %d days\n", atoi(parts[5]))
		}
		out("},\n")
	}
	out(ftr)

	check(os.WriteFile(os.Args[2], output.Bytes(), 0644))
}

func out(format string, args ...interface{}) {
	fmt.Fprintf(&output, format, args...)
}

func month(s string) time.Month {
	for nr, name := range []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	} {
		if s == name {
			return time.Month(nr + 1)
		}
	}
	check(fmt.Errorf("cannot decode %q into a month", s))
	return 0
}

func atoi(s string) int {
	r, err := strconv.Atoi(s)
	check(err)
	return r
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

/* 
ConversionTable defines matching Julian and Gregorian dates. It is consulted by, e.g., Date.Convert(). The contained ConversionEntry's won't normally be of use outside of this module, but they are coded as exportable (as uppercase symbols) so they can be inspected.

This table reflects https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars, original source: the Nautical almanac of the United Kingdom and United States (1961). This table however knows that year zero doesn't exist (we go 2BC, 1BC, 1AD, 2AD), therefore, years before 0 are generated one-off relative to the above source refrence.
*/
var ConversionTable = [...]ConversionEntry{
	// [...] is syntactic sugar to let the compiler figure out the array size. That way
	// we get a fixed size array and not a slice.
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

	for _, line := range strings.Split(string(content), "\n") {
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

		// Adjust for BC: there is no year zero.
		if yr <= 0 {
			yr--
		}

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

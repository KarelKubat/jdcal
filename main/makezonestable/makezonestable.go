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

	"github.com/KarelKubat/jdcal"
)

const (
	hdr = `package jdcal

// Generated file, don't edit by hand. Look for a nearby Makefile.

import (
	"time"
)

/*
ZonesTable defines matching geographical zones to the date of adoption of the
Gregorian calendar. Some zones temporarily reverted to the Julian calendar, these
have multiple cutover dates.

This table reflects https://en.wikipedia.org/wiki/List_of_adoption_dates_of_the_Gregorian_calendar_by_country.


Note: The Cutovers entries are the dates where that period STOPPED. So: "{Year: 1918, Month: time.April, Day: 17, Type: Julian}" is the Julian date 1918/04/16, when the Julian calendar was abandoned in favor of "the other one", being the Gregorian.

Another note: All symbols of the table are exported (upper case). You can skip reading the table; rather, have a look at the "Zones*()" functions that can do the lifting.
*/
var ZonesTable = [...]ZoneEntry{
	// [...] is syntactic sugar to let the compiler figure out the array size. That way
	// we get a fixed size array and not a slice.	
`
	ftr = `
}
`
)

var output bytes.Buffer

type zoneInfo struct {
	name     string
	cutovers []string
}

func main() {
	if len(os.Args) != 3 {
		check(errors.New("usage: go run what/ever/makezonestable.go INPUT OUTPUT"))
	}
	content, err := os.ReadFile(os.Args[1])
	check(err)

	out(hdr)
	splitRe := regexp.MustCompile(`[ \t]+`)
	emptyLine := regexp.MustCompile(`^\s*$`)
	curr := zoneInfo{}

	lineno := 0
	for _, line := range strings.Split(string(content), "\n") {
		lineno++

		if strings.HasPrefix(line, "//") || emptyLine.MatchString(line) {
			continue
		}
		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
			flush(curr)
			line = strings.TrimSpace(line)
			curr.name = line
			curr.cutovers = nil
			continue
		}
		line = strings.TrimSpace(line)
		parts := splitRe.Split(line, -1)
		var tp jdcal.Type
		switch parts[0] {
		case "to-gd":
			tp = jdcal.Julian
		case "to-jd":
			tp = jdcal.Gregorian
		default:
			check(fmt.Errorf("line %d: %q needs either to-gd or to-jd",
				lineno, line))
		}

		if parts[1] == "FIRST" {
			curr.cutovers = append(curr.cutovers, fmt.Sprintf("First(%v)", tp))
		} else {
			nrs := strings.Split(parts[1], "/")
			if len(nrs) != 3 {
				check(fmt.Errorf("line %d: %q has a malformed date, not 3 parts but %d",
					lineno, line, len(nrs)))
			}
			dt, err := jdcal.NewDate(
				jdcal.Year(atoi(nrs[0])),
				time.Month(atoi(nrs[1])),
				atoi(nrs[2]),
				tp)
			check(err)
			curr.cutovers = append(curr.cutovers,
				fmt.Sprintf("{Year:%v, Month:time.%v, Day:%v, Type:%v}",
					dt.Year, dt.Month, dt.Day, dt.Type))
		}
	}
	flush(curr)

	out(ftr)
	check(os.WriteFile(os.Args[2], output.Bytes(), 0644))
}

func flush(z zoneInfo) {
	if z.name == "" {
		if len(z.cutovers) > 0 {
			check(fmt.Errorf("%v has no name, but cutovers", z.name))
		}
		return
	}
	out("  {\n    Name: %q,\n    Cutovers: []Date{\n", z.name)
	for _, c := range z.cutovers {
		out("    %v,\n", c)
	}
	out("    },\n},\n")
}

func out(format string, args ...interface{}) {
	fmt.Fprintf(&output, format, args...)
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

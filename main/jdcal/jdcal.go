package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KarelKubat/flagnames"
	"github.com/KarelKubat/jdcal"
)

var (
	flagJulian    = flag.Bool("julian", false, "when true, commandline date is taken as a Julian date")
	flagGregorian = flag.Bool("gregorian", false, "when true, commandline date is taken as a Gregorian date")
)

const (
	usageInfo = `
Usage: jdcal [-j] [-g] YYYY/MM/DD [YYYY-MM-DD...]
Converts the given date(s) either from Julian to Gregorian (when -j is given),
or from Gregorian to Julian (when -g is given).
`
)

func main() {
	flagnames.Patch()
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usageInfo)
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagJulian == *flagGregorian {
		check(errors.New("one of -j or -g must be given, but not both"))
	}
	if flag.NArg() == 0 {
		flag.Usage()
	}

	var tp jdcal.Type
	if *flagJulian {
		tp = jdcal.Julian
	} else {
		tp = jdcal.Gregorian
	}

	for _, arg := range flag.Args() {
		parts := strings.Split(arg, "/")
		if len(parts) != 3 {
			check(fmt.Errorf("malformed argument %q, want YYYY/MM/DD", arg))
		}
		dt, err := jdcal.New(jdcal.Date{
			Year:  atoi(arg, parts[0]),
			Month: time.Month(atoi(arg, parts[1])),
			Day:   atoi(arg, parts[2]),
			Type:  tp,
		})
		check(err)
		ot, err := dt.Convert()
		check(err)
		fmt.Println(dt, "is", ot)
	}
}

func atoi(fullArg, arg string) int {
	r, err := strconv.Atoi(arg)
	if err != nil {
		check(fmt.Errorf("part %q of %q is not a number: %v", arg, fullArg, err))
	}
	return r
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

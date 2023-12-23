package main

import (
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
or from Gregorian to Julian (when -g is given). Either -j, or -g, or both must
be given.
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
	if len(os.Args) < 2 {
		flag.Usage()
	}
	if flag.NArg() == 0 || (!*flagJulian && !*flagGregorian) {
		flag.Usage()
	}

	types := map[jdcal.Type]bool{
		jdcal.Gregorian: *flagGregorian,
		jdcal.Julian:    *flagJulian,
	}

	for _, arg := range flag.Args() {
		parts := strings.Split(arg, "/")
		if len(parts) != 3 {
			check(fmt.Errorf("malformed argument %q, want YYYY/MM/DD", arg))
		}
		for tp, act := range types {
			if !act {
				continue
			}
			dt, err := jdcal.New(atoi(arg, parts[0]), time.Month(atoi(arg, parts[1])),
				atoi(arg, parts[2]), tp)
			check(err)
			ot, err := dt.Convert()
			check(err)
			fmt.Println(dt, "is", ot)
		}
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

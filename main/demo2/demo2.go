package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	// Advancing dates
	// ---------------
	jd, err := jdcal.New(300, time.February, 27, jdcal.Julian)
	check(err)

	for i := 0; i < 6; i++ {
		gd, err := jd.Convert()
		check(err)
		fmt.Println(jd, "is", gd)
		jd = jd.Advance()
	}

	// Output:
	// Julian 0300/02/27 is Gregorian 0300/02/27
	// Julian 0300/02/28 is Gregorian 0300/02/28
	// Julian 0300/02/29 is Gregorian 0300/03/01  # BOOM, Julian is a day wrong
	// Julian 0300/03/01 is Gregorian 0300/03/02
	// Julian 0300/03/02 is Gregorian 0300/03/03
	// Julian 0300/03/03 is Gregorian 0300/03/04

	// Note that the Julian calendar knows a February 29th, the Gregorian one doesn't.
	// The two calendars diverge after February 28th. This is historically correct.

	// Testing leap years
	// ------------------
	for _, yr := range []jdcal.Year{1886, 1800, 2000} {
		for _, tp := range []jdcal.Type{jdcal.Julian, jdcal.Gregorian} {
			cyr := jdcal.CalendarYear{Year: yr, Type: tp}
			fmt.Println(cyr, "is a leap year:", cyr.IsLeap())
		}
	}

	// Output:

	// Julian 1886 is a leap year: false       # 1886: Julian and Gregorian IsLeap agree
	// Gregorian 1886 is a leap year: false

	// Julian 1800 is a leap year: true        # 1800: Julian and Gregorian IsLeap disagree
	// Gregorian 1800 is a leap year: false

	// Julian 2000 is a leap year: true        # 2000: Julian and Gregorian agree
	// Gregorian 2000 is a leap year: true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

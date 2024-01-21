// main/demo2/demo2.go
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
	jd, err := jdcal.NewDate(300, time.February, 27, jdcal.Julian)
	check(err)

	for i := 0; i < 6; i++ {
		gd, err := jd.Convert()
		check(err)
		fmt.Println(jd, "is", gd)
		jd = jd.Forward()
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

	// Going backward
	// --------------
	for i := 0; i < 6; i++ {
		gd, err := jd.Convert()
		check(err)
		fmt.Println(jd, "is", gd)
		jd = jd.Backward()
	}

	// Output:
	// Julian 0300/03/04 is Gregorian 0300/03/05
	// Julian 0300/03/03 is Gregorian 0300/03/04
	// Julian 0300/03/02 is Gregorian 0300/03/03
	// Julian 0300/03/01 is Gregorian 0300/03/02
	// Julian 0300/02/29 is Gregorian 0300/03/01
	// Julian 0300/02/28 is Gregorian 0300/02/28

	// Testing leap years
	// ------------------
	for _, yr := range []jdcal.Year{1796, 1797, 1800, 2000} {
		for _, tp := range []jdcal.Type{jdcal.Julian, jdcal.Gregorian} {
			cyr, err := jdcal.NewCalendarYear(yr, tp)
			check(err)
			fmt.Println(cyr, "is a leap year:", cyr.IsLeap())
		}
	}

	// Output:

	// Julian 1796 is a leap year: true        # Standard leap year (divisible by 4)
	// Gregorian 1796 is a leap year: true     # or standard non-leap: Julian and Gregorian
	// Julian 1797 is a leap year: false       # agree
	// Gregorian 1797 is a leap year: false

	// Julian 1800 is a leap year: true        # Century: IsLeap disagrees
	// Gregorian 1800 is a leap year: false

	// Julian 2000 is a leap year: true        # Millenium: Julian and Gregorian agree
	// Gregorian 2000 is a leap year: true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

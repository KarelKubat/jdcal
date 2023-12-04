/*
Package jdcal provides conversions from dates on the Julian calendar to the Gregorian calendar and vice versa. Furthermore helper functions are provided for small tasks, such as: date validation, is a year a leap year (which depends on the calendar type), comparisons between dates (before/after/equal), advancing a date while honoring leap years.

Example 1: Conversions

	package main

	import (
		"fmt"
		"log"
		"time"

		"github.com/KarelKubat/jdcal"
	)

	func main() {
		// October 5th (Julian) was the Papal announcement to skip 10 days.
		// The new date would be October 15 (Gregorian).
		jd0, err := jdcal.New(1582, time.October, 5, jdcal.Julian)
		check(err)

		// to Gregorian
		gd, err := jd0.Convert()
		check(err)
		fmt.Println("From Julian to Gregorian:", jd0, "is", gd)

		// back to Julian
		jd1, err := gd.Convert()
		check(err)
		fmt.Println("And back again:", gd, "is", jd1)

		// Output:
		// From Julian to Gregorian: Julian 1582/10/05 is Gregorian 1582/10/15
		// And back again: Gregorian 1582/10/15 is Julian 1582/10/05
	}

	func check(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

Example 2: Leap years on Julian calendars are different

	package main

	import (
		"fmt"
		"log"
		"time"

		"github.com/KarelKubat/jdcal"
	)

	func main() {
		jd, err := jdcal.New(300, time.February, 26, jdcal.Julian)
		check(err)

		for i := 0; i < 10; i++ {
			gd, err := jd.Convert()
			check(err)
			fmt.Println(jd, "is", gd)
			jd = jd.Advance()
		}

		// Output:
		// Julian 0300/02/26 is Gregorian 0300/02/26
		// Julian 0300/02/27 is Gregorian 0300/02/27
		// Julian 0300/02/28 is Gregorian 0300/02/28
		// Julian 0300/02/29 is Gregorian 0300/03/01
		// Julian 0300/03/01 is Gregorian 0300/03/02
		// Julian 0300/03/02 is Gregorian 0300/03/03
		// Julian 0300/03/03 is Gregorian 0300/03/04
		// Julian 0300/03/04 is Gregorian 0300/03/05
		// Julian 0300/03/05 is Gregorian 0300/03/06
		// Julian 0300/03/06 is Gregorian 0300/03/07

		// Note that the Julian calendar knows a February 29th, the Gregorian one doesn't.
		// The two calendars diverge after February 28th. This is historically correct;
		// before the date the two calendars would be in sync and in the coming centuries
		// they would diverge even more.
	}

	func check(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
*/
package jdcal

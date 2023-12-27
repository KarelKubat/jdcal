/*
Package jdcal provides conversions from dates on the Julian calendar to the Gregorian calendar and vice versa, and knows about zones (geographical locations) where such switches occurred, and at which dates.

Furthermore helper functions are provided for small tasks, such as: date validation, is a year a leap year (which depends on the calendar type), comparisons between dates (before/after/equal), advancing a date while honoring leap years.

The distribution also contains a CLI main/jdcal/jdcal.go which provides a command line interface.

CLI Example 1: Conversions

	# after `go install main/jdcal.go`:
	jdcal convert 1600/01/01 --julian  # assume the input date is Julian
	Julian 1600/01/01 is Gregorian 1600/01/11

CLI Example 2: Known zones

	jdcal zones --match netherlands  # --match restricts the output
	Belgium (Southern Netherlands)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1582/12/20
	Netherlands (Brabant)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1582/12/14
	Netherlands (Drenthe)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1701/04/30
	Netherlands (Frisia)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1701/12/31
	Netherlands (Gelderland)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1700/06/12
	Netherlands (Groningen City)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1583/01/01
	  Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	  Switched to   the Gregorian calendar   on   Julian 1700/12/31
	Netherlands (Holland)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1583/01/01
	Netherlands (Utrecht, Overijssel)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1700/11/30
	Netherlands (Zeeland, States General)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1582/12/14

API Example 1: Conversions

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

API Example 2: Leap years on Julian calendars are different

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

API Example 3: Zones

	package main

	import (
		"fmt"

		"github.com/KarelKubat/jdcal"
	)

	func main() {
		for _, e := range jdcal.ZonesByName("netherlands") {
			fmt.Println(e)
		}

		// Output:
		//   Belgium (Southern Netherlands)
		// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
		// 		Switched to   the Gregorian calendar   on   Julian 1582/12/20
		// ... and so on
	}
*/
package jdcal

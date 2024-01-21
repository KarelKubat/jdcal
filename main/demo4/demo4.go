// main/demo4/demo4.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

const (
	zoneName = "Netherlands (Groningen City)"
)

func main() {
	zone, err := jdcal.SingleZone(zoneName)
	check(err)
	fmt.Println(zone)
	// Netherlands (Groningen City)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1583/01/01
	//   Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	//   Switched to   the Gregorian calendar   on   Julian 1700/12/31

	// Some dates that lie in between the cutovers.
	for _, year := range []jdcal.Year{1580, 1590, 1600, 1800} {
		test(year, time.January, 1, zone)
	}

	// Output:
	//   1580/01/01 is a Julian date
	//   1590/01/01 is a Gregorian date
	//   1600/01/01 is a Julian date
	//   1800/01/01 is a Gregorian date

	// Just around the exact cutover dates.
	for _, date := range []struct {
		year  jdcal.Year
		month time.Month
		day   int
	}{
		// Around the first switch from Julian into the Gregorian calendar
		{year: 1582, month: time.December, day: 31},
		{year: 1583, month: time.January, day: 1},
		{year: 1583, month: time.January, day: 2},
		// Output:
		//   1582/12/31 is a Julian date
		//   1583/01/01 is a Julian date
		//   1583/01/02 is neither a Julian nor a Gregorian date

		// Around the second switch from Gregorian back to Julian
		{year: 1594, month: time.November, day: 9},
		{year: 1594, month: time.November, day: 10},
		{year: 1594, month: time.November, day: 11},
		// Output:
		//   1594/11/09 can be both a Julian and a Gregorian date
		//   1594/11/10 can be both a Julian and a Gregorian date
		//   1594/11/11 is a Julian date

		// Around the third switch back to Gregorian
		{year: 1700, month: time.December, day: 30},
		{year: 1700, month: time.December, day: 31},
		{year: 1701, month: time.January, day: 1},
		// Output:
		//   1700/12/30 is a Julian date
		//   1700/12/31 is a Julian date
		//   1701/01/01 is neither a Julian nor a Gregorian date
	} {
		test(date.year, date.month, date.day, zone)
	}
}

func test(year jdcal.Year, month time.Month, day int, z jdcal.ZoneEntry) {
	d, err := jdcal.New(year, month, day, jdcal.Julian)
	check(err)
	jdInZone, err := d.InZone(z)
	check(err)

	d, err = jdcal.New(year, month, day, jdcal.Gregorian)
	check(err)
	gdInZone, err := d.InZone(z)
	check(err)

	fmt.Printf("%4.4d/%2.2d/%2.2d ", year, int(month), day)
	switch {
	case jdInZone && gdInZone:
		fmt.Println("can be both a Julian and a Gregorian date")
	case !jdInZone && !gdInZone:
		fmt.Println("is neither a Julian nor a Gregorian date")
	case jdInZone:
		fmt.Println("is a Julian date")
	default:
		fmt.Println("is a Gregorian date")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

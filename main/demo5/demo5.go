// main/demo5/demo5.go
package main

import (
	"fmt"
	"log"

	"github.com/KarelKubat/jdcal"
)

func main() {

	for _, yr := range []jdcal.Year{1369, 1370, 1371} {
		cyr, err := jdcal.NewCalendarYear(yr, jdcal.Julian)
		check(err)

		for h := jdcal.AshWednesday; h <= jdcal.Pentecost; h++ {
			dt, err := cyr.HolidayDate(h)
			check(err)
			wd, err := dt.Weekday()
			check(err)
			fmt.Println(h, "in", cyr, "falls on", wd, dt)
		}
		fmt.Println()
	}

	// Output:
	// Ash Wednesday in Julian 1369 falls on Wednesday Julian 1369/02/07
	// Good Friday in Julian 1369 falls on Friday Julian 1369/03/23
	// Easter in Julian 1369 falls on Sunday Julian 1369/03/25
	// Ascension Day in Julian 1369 falls on Thursday Julian 1369/05/03
	// Pentecost in Julian 1369 falls on Sunday Julian 1369/05/13

	// Ash Wednesday in Julian 1370 falls on Wednesday Julian 1370/02/27
	// Good Friday in Julian 1370 falls on Friday Julian 1370/04/12
	// Easter in Julian 1370 falls on Sunday Julian 1370/04/14
	// Ascension Day in Julian 1370 falls on Thursday Julian 1370/05/23
	// Pentecost in Julian 1370 falls on Sunday Julian 1370/06/02

	// Ash Wednesday in Julian 1371 falls on Wednesday Julian 1371/02/19
	// Good Friday in Julian 1371 falls on Friday Julian 1371/04/04
	// Easter in Julian 1371 falls on Sunday Julian 1371/04/06
	// Ascension Day in Julian 1371 falls on Thursday Julian 1371/05/15
	// Pentecost in Julian 1371 falls on Sunday Julian 1371/05/25
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/KarelKubat/jdcal"
)

func main() {

	for _, yr := range []jdcal.Year{1369, 1370, 1371} {
		cyr := jdcal.CalendarYear{Year: yr, Type: jdcal.Julian}

		for _, h := range []jdcal.Holiday{
			jdcal.Easter,
			jdcal.Ascension,
			jdcal.Pentecost,
		} {
			dt, err := cyr.HolidayDate(h)
			check(err)
			wd, err := dt.Weekday()
			check(err)
			fmt.Println(h, "in", cyr, "falls on", wd, dt)
		}
		fmt.Println()
	}

	// Output:
	// Easter in Julian 1369 falls on Sunday Julian 1369/03/25
	// Ascension day in Julian 1369 falls on Thursday Julian 1369/05/03
	// Pentecost in Julian 1369 falls on Sunday Julian 1369/05/13

	// Easter in Julian 1370 falls on Sunday Julian 1370/04/14
	// Ascension day in Julian 1370 falls on Thursday Julian 1370/05/23
	// Pentecost in Julian 1370 falls on Sunday Julian 1370/06/02

	// Easter in Julian 1371 falls on Sunday Julian 1371/04/06
	// Ascension day in Julian 1371 falls on Thursday Julian 1371/05/15
	// Pentecost in Julian 1371 falls on Sunday Julian 1371/05/25
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

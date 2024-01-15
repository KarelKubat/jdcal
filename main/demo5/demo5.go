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
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

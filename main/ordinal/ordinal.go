package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	for _, ymd := range []jdcal.YMD{
		// Start of epoch
		// {Year: -500, Month: time.January, Day: 1},

		// Around negative leap
		{Year: -301, Month: time.February, Day: 28},
		{Year: -301, Month: time.March, Day: 1},

		// Aroound positive leap
		{Year: 300, Month: time.February, Day: 28},
		{Year: 300, Month: time.March, Day: 1},

		// My fav
		{Year: 1962, Month: time.August, Day: 19},
	} {
		for _, tp := range []jdcal.Type{jdcal.Julian, jdcal.Gregorian} {
			dt, err := jdcal.New(ymd.Year, ymd.Month, ymd.Day, tp)
			check(err)
			fmt.Printf("%-22v --> ", dt)
			ord, err := dt.Ordinal()
			check(err)
			fmt.Printf("%6d --> ", ord)
			back, err := jdcal.OrdinalToDate(ord, dt.Type)
			check(err)
			fmt.Println(back)

			eq, err := dt.Equal(back)
			check(err)
			if !eq {
				check(fmt.Errorf("%v and %v mismatch", dt, back))
			}

		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

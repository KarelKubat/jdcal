package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	for _, ymd := range []jdcal.YMD{
		// Near start of epoch
		{Year: -500, Month: time.January, Day: 1},

		// Around negative leap
		{Year: -301, Month: time.February, Day: 28},
		{Year: -301, Month: time.March, Day: 1},

		// Aroound positive leap
		{Year: 300, Month: time.February, Day: 28},
		{Year: 300, Month: time.March, Day: 1},

		// Somewhere in the 20th century
		{Year: 1962, Month: time.August, Day: 19},

		// Near end of epoch
		{Year: 2100, Month: time.January, Day: 1},
	} {
		for _, tp := range []jdcal.Type{jdcal.Gregorian, jdcal.Julian} {
			dt, err := jdcal.New(ymd.Year, ymd.Month, ymd.Day, tp)
			check(err)
			ord, err := dt.Ordinal()
			check(err)
			back, err := ord.Date(dt.Type)
			check(err)

			fmt.Printf("%-22v --> %6d --> %v\n", dt, ord, back)

			eq, err := dt.Equal(back)
			check(err)
			if !eq {
				check(fmt.Errorf("%v and %v mismatch", dt, back))
			}

		}
	}

	// Output (actual oridnal number may differ):
	// Gregorian -0500/01/01  -->   2191 --> Gregorian -0500/01/01
	// Julian -0500/01/01     -->   2186 --> Julian -0500/01/01
	// Gregorian -0301/02/28  -->  74933 --> Gregorian -0301/02/28
	// Julian -0301/02/28     -->  74928 --> Julian -0301/02/28
	// Gregorian -0301/03/01  -->  74934 --> Gregorian -0301/03/01
	// Julian -0301/03/01     -->  74930 --> Julian -0301/03/01
	// Gregorian 0300/02/28   --> 294444 --> Gregorian 0300/02/28
	// Julian 0300/02/28      --> 294444 --> Julian 0300/02/28
	// Gregorian 0300/03/01   --> 294445 --> Gregorian 0300/03/01
	// Julian 0300/03/01      --> 294446 --> Julian 0300/03/01
	// Gregorian 1962/08/19   --> 901649 --> Gregorian 1962/08/19
	// Julian 1962/08/19      --> 901662 --> Julian 1962/08/19
	// Gregorian 2100/01/01   --> 951823 --> Gregorian 2100/01/01
	// Julian 2100/01/01      --> 951836 --> Julian 2100/01/01
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// main/demo6/demo6.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	jyr, err := jdcal.NewCalendarYear(1900, jdcal.Julian)
	check(err)
	gyr, err := jdcal.NewCalendarYear(1900, jdcal.Gregorian)
	check(err)

	jdpm := jyr.DaysPerMonth()
	gdpm := gyr.DaysPerMonth()

	for m := time.January; m <= time.December; m++ {
		fmt.Printf("%-10s: in Julian %2.2d, in Gregorian %2.2d days\n", m, jdpm[m], gdpm[m])
	}

	// Output:
	// January   : in Julian 31, in Gregorian 31 days
	// February  : in Julian 29, in Gregorian 28 days
	// March     : in Julian 31, in Gregorian 31 days
	// April     : in Julian 30, in Gregorian 30 days
	// May       : in Julian 31, in Gregorian 31 days
	// June      : in Julian 30, in Gregorian 30 days
	// July      : in Julian 31, in Gregorian 31 days
	// August    : in Julian 31, in Gregorian 31 days
	// September : in Julian 30, in Gregorian 30 days
	// October   : in Julian 31, in Gregorian 31 days
	// November  : in Julian 30, in Gregorian 30 days
	// December  : in Julian 31, in Gregorian 31 days
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// main/demo1/demo1.go
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
	jd0, err := jdcal.NewDate(1582, time.October, 5, jdcal.Julian)
	check(err)

	// To switch to a slower, but longer tested algorithm:
	// jdcal.ConvertByLookup()

	// to Gregorian
	gd, err := jd0.Convert()
	check(err)
	wd, err := gd.Weekday()
	check(err)
	fmt.Println("From Julian to Gregorian:", jd0, "is", gd, "and it's a", wd)

	// back to Julian
	jd1, err := gd.Convert()
	check(err)
	wd, err = jd1.Weekday()
	check(err)
	fmt.Println("And back again:", gd, "is", jd1, "and it's a", wd)

	// Output:
	// 	From Julian to Gregorian: Julian 1582/10/05 is Gregorian 1582/10/15 and it's a Friday
	// 	And back again: Gregorian 1582/10/15 is Julian 1582/10/05 and it's a Friday
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

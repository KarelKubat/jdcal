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

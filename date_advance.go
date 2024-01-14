package jdcal

import (
	"time"
)

/*
Advance returns a date that is one day later, honoring leap days for Julian or Gregorian calendars. The reference date (receiver) is not modified.
Example:

	jd0, err := jdcal.Date(1900, time.February, 28, jdcal.Julian)
	if err != nil {...}
	jd1 := jd0.Advance()  // Copy of jd but a day ahead
	fmt.Println(jd1)      // February 29th

	gd, err := jdcal.Date(1900, time.February, 28, jdcal.Gregorian)
	if err != nil {...}
	gd = gd.Advance()     // advance gd itself
	fmt.Println(gd)       // March 1st
*/
func (d Date) Advance() Date {
	// Nr of month days, with 28 for February (handled separately).
	daysPerMonth := []int{
		0, // filler, January is #1
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	ret := d

	if ret.Day < daysPerMonth[ret.Month] {
		ret.Day++
		return ret
	}
	if ret.Month == time.February {
		if ret.Day == 28 && ret.Year.IsLeap(ret.Type) {
			ret.Day++
		} else {
			ret.Month = time.March
			ret.Day = 1
		}
		return ret
	}
	if ret.Month == time.December {
		// There is no year zero, we go from 1BC to 1AD.
		if ret.Year == -1 {
			ret.Year += 2
		} else {
			ret.Year++
		}
		ret.Day = 1
		ret.Month = time.January
		return ret
	}
	ret.Day = 1
	ret.Month++
	return ret
}

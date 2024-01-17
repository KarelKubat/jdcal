package jdcal

import (
	"time"
)

/*
Forward returns a date that is one day later, honoring leap days for Julian or Gregorian calendars. The reference date (receiver) is not modified.
Example:

	jd0, err := jdcal.Date(1900, time.February, 28, jdcal.Julian)
	if err != nil {...}
	jd1 := jd0.Forward()  // Copy of jd but a day ahead
	fmt.Println(jd1)      // February 29th

	gd, err := jdcal.Date(1900, time.February, 28, jdcal.Gregorian)
	if err != nil {...}
	gd = gd.Forward()     // Advance gd itself
	fmt.Println(gd)       // March 1st
*/
func (d Date) Forward() Date {
	// Nr of month days, with 28 for February (handled separately).
	daysPerMonth := CalendarYear{Year: d.Year, Type: d.Type}.DaysPerMonth()
	ret := d

	// Simple increase within 1 month
	if ret.Day < daysPerMonth[ret.Month] {
		ret.Day++
		return ret
	}

	// Going to the 1st day of the next month
	ret.Day = 1
	if ret.Month < time.December {
		ret.Month++
		return ret
	}
	// December goes to January. Also, there is no year zero, we go from 1BC to 1AD.
	ret.Month = time.January
	if ret.Year == -1 {
		ret.Year += 2
	} else {
		ret.Year++
	}
	return ret
}

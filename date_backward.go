package jdcal

import (
	"fmt"
	"time"
)

/*
Backward returns a date that is one day before, hononing leap years for Julian or Gregorian calendars. The reference date (receiver) is not modified.
Example:

	jd0, err := jdcal.Date(1900, time.March, 1, jdcal.Julian)
	if err != nil {...}
	jd1 := jd0.Backward()  // Copy of jd but a day before
	fmt.Println(jd1)       // February 29th

	gd, err := jdcal.Date(1900, time.March, 1, jdcal.Gregorian)
	if err != nil {...}
	gd = gd.Backward()    // Decrease gd itself
	fmt.Println(gd)       // February 28th
*/
func (d Date) Backward() Date {
	cyr, err := NewCalendarYear(d.Year, d.Type)
	if err != nil {
		panic(fmt.Sprintf("internal error: Date.Backeward failed to construct a CalendarYear: %v", err))
	}
	daysPerMonth := cyr.DaysPerMonth()
	ret := d

	// Within a month
	if ret.Day > 1 {
		ret.Day--
		return ret
	}

	// February or later: decrease the month
	if ret.Month > time.January {
		ret.Month--
		ret.Day = daysPerMonth[ret.Month]
		return ret
	}

	// January 1st; go to December of the previous year. Adjust for BC (-1 is year zero)
	ret.Month = time.December
	ret.Day = daysPerMonth[ret.Month]
	if ret.Year == 1 {
		ret.Year = -1
	} else {
		ret.Year--
	}
	return ret
}

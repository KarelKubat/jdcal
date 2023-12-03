package jdcal

/*
https://en.wikipedia.org/wiki/Leap_year:

The historic Julian calendar has three common years of 365 days followed by a leap year of 366 days, by extending February to 29 days rather than the common 28.

The Gregorian calendar, the world's most widely used civil calendar, makes a further adjustment for the small error in the Julian algorithm. Again each leap year has 366 days instead of 365. This extra leap day occurs in each year that is an integer multiple of 4 (except for years evenly divisible by 100, but not by 400).
*/

func (d *Date) IsLeap() bool {
	if d.Year%4 != 0 {
		return false
	}
	if d.Type == Julian {
		return true
	}
	return d.Year%100 != 0 || d.Year%400 == 0
}

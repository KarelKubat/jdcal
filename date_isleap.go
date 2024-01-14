package jdcal

/*
IsLeap is true when a date's year indicate that that year should be a leap year. The month and day
don't matter, but the type does (Julian or Gregorian). IsLeap implements the following definition:

https://en.wikipedia.org/wiki/Leap_year:

The historic Julian calendar has three common years of 365 days followed by a leap year of 366 days, by extending February to 29 days rather than the common 28.

The Gregorian calendar, the world's most widely used civil calendar, makes a further adjustment for the small error in the Julian algorithm. Again each leap year has 366 days instead of 365. This extra leap day occurs in each year that is an integer multiple of 4 (except for years evenly divisible by 100, but not by 400).

Example:

	jd, err := jdcal.New(1900, time.June, 12, jdcal.Julian)
	if err != nil {...}
	fmt.Println(jd.IsLeap())  // true

	gd, err := jdcal.New(1900, time.June, 12, jdcal.Gregorian)
	if err != nil {...}
	fmt.Println(gd.IsLeap())  // false
*/
func (d Date) IsLeap() bool {
	year := d.Year

	// Adjust negative years, since there is no year zero. So year -1 is a leap year, -5 too, etc.
	if year < 0 {
		year++
	}

	if year%4 != 0 {
		return false
	}
	if d.Type == Julian {
		return true
	}
	return year%100 != 0 || year%400 == 0
}

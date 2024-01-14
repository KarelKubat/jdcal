package jdcal

/*
IsLeap is true when a year indicate that that year should be a leap year for the given calendar type. IsLeap implements the following definition:

https://en.wikipedia.org/wiki/Leap_year:

The historic Julian calendar has three common years of 365 days followed by a leap year of 366 days, by extending February to 29 days rather than the common 28.

The Gregorian calendar, the world's most widely used civil calendar, makes a further adjustment for the small error in the Julian algorithm. Again each leap year has 366 days instead of 365. This extra leap day occurs in each year that is an integer multiple of 4 (except for years evenly divisible by 100, but not by 400).

Example:

	var cyr jdcal.CalendarYear

	cyr = {
		Year: 1900,
		Type: jdcal.Julian,
	}
	fmt.Println(cyr.IsLeap())  // true
	cyr.Type = jdcal.Gregorian
	fmt.Println(cyr.IsLeap())  // false
*/
func (cYear CalendarYear) IsLeap() bool {
	// Adjust negative years, since there is no year zero. So year -1 is a leap year, -5 too, etc.
	yr := cYear.Year
	tp := cYear.Type

	if yr < 0 {
		yr++
	}

	if yr%4 != 0 {
		return false
	}
	if tp == Julian {
		return true
	}
	return yr%100 != 0 || yr%400 == 0

}

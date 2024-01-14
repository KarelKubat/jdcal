package jdcal

/*
IsLeap is true when a year indicate that that year should be a leap year for a given calendar type. IsLeap implements the following definition:

https://en.wikipedia.org/wiki/Leap_year:

The historic Julian calendar has three common years of 365 days followed by a leap year of 366 days, by extending February to 29 days rather than the common 28.

The Gregorian calendar, the world's most widely used civil calendar, makes a further adjustment for the small error in the Julian algorithm. Again each leap year has 366 days instead of 365. This extra leap day occurs in each year that is an integer multiple of 4 (except for years evenly divisible by 100, but not by 400).

Example:

	var yr jdcal.Year

	yr = 1900
	fmt.Println(yr.IsLeap(jdcal.Julian))     // true
	fmt.Println(yr.IsLeap(jdcal.Gregorian))  // false
*/
func (year Year) IsLeap(tp Type) bool {
	// Adjust negative years, since there is no year zero. So year -1 is a leap year, -5 too, etc.
	if year < 0 {
		year++
	}

	if year%4 != 0 {
		return false
	}
	if tp == Julian {
		return true
	}
	return year%100 != 0 || year%400 == 0

}

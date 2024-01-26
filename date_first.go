package jdcal

import (
	"time"
)

/*
First returns the first convertible day for a given type. Dates before this cannot be converted; Convert() would throw an error. This is a limitation of the ConversionTable. Example:

	jd := jdcal.First(jdcal.Julian)  // First convertible date
	gd, err := jd.Convert()          // err will be nil, jd can be converted

The first convertible dates point to the same day, despite being different day numbers on the Julian or Gregorian calendars.
*/
func First(tp Type) Date {
	switch Algorithm {
	case ByProgression:
		return Date{Year: StartProgressionYear, Month: time.January, Day: 1, Type: tp}
	case ByLookup:
		if tp == Julian {
			return ConversionTable[0].JDate
		}
		return ConversionTable[0].GDate
	default:
		panic("internal error: algorithm mismatch in Date.First")
	}
}

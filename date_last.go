package jdcal

import "time"

/*
Last returns the last convertible date for a given type. Dates after this cannot be converted;
Convert() would throw an error. This is a limitation of the ConversionTable. Example:

	gd := jdcal.Last(jdcal.Gregorian)  // Last convertible date
	gd = gd.Forward()                  // Move 1 day forward
	jd, err := gd.Convert()            // err will be set, gd cannot be converted
*/
func Last(tp Type) Date {
	switch Algorithm {
	case ByProgression:
		return Date{Year: EndProgressionYear, Month: time.January, Day: 1, Type: tp}
	case ByLookup:
		if tp == Julian {
			return ConversionTable[len(ConversionTable)-1].JDate
		}
		return ConversionTable[len(ConversionTable)-1].GDate
	default:
		panic("internal error: algorithm mismatch in Date.Last")
	}
}

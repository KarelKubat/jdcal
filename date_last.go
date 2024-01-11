package jdcal

/*
Last returns the last convertible date for a given type. Dates after this cannot be converted;
Convert() would throw an error. This is a limitation of the ConversionTable. Example:

	gd := jdcal.Last(jdcal.Gregorian)  // Last convertible date
	gd = gd.Advance()                  // Move 1 day forward
	jd, err := gd.Convert()            // err will be set, gd cannot be converted
*/
func Last(dt Type) Date {
	if dt == Julian {
		return ConversionTable[len(ConversionTable)-1].JDate
	}
	return ConversionTable[len(ConversionTable)-1].GDate
}

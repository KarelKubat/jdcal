package jdcal

/*
First returns the first convertible day for a given type. Dates before this cannot be converted;
Convert() would throw an error. This is a limitation of the ConversionTable. Example:

	jd := jdcal.First(jdcal.Julian)  // First convertible date
	gd, err := jd.Convert()          // err will be nil, jd can be converted
*/
func First(dt Type) Date {
	if dt == Julian {
		return ConversionTable[0].JDate
	}
	return ConversionTable[0].GDate
}

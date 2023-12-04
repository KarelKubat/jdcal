package jdcal

func First(dt Type) Date {
	if dt == Julian {
		return ConversionTable[0].JDate
	}
	return ConversionTable[0].GDate
}

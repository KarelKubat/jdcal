package jdcal

func Last(dt Type) Date {
	if dt == Julian {
		return ConversionTable[len(ConversionTable)-1].JDate
	}
	return ConversionTable[len(ConversionTable)-1].GDate
}

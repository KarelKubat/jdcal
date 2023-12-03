package jdcal

func First(dt Type) Date {
	if dt == Julian {
		return Table[0].JDate
	}
	return Table[0].GDate
}

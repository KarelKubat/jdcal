package jdcal

func Last(dt Type) Date {
	if dt == Julian {
		return Table[len(Table)-1].JDate
	}
	return Table[len(Table)-1].GDate
}

package jdcal

func (t Type) String() string {
	return []string{"Gregorian", "Julian"}[t]
}

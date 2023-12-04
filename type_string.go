package jdcal

/*
String returns "Gregorian" for jdcal.Gregorian or "Julian" for jdcal.Julian.
*/
func (t Type) String() string {
	return []string{"Gregorian", "Julian"}[t]
}

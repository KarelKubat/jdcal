package jdcal

/*
String returns the string representation for a Holiday: "Easter", "Ascension day", etc.
*/
func (h Holiday) String() string {
	return []string{
		"", // start sentinel
		"Ash Wednesday",
		"Good Friday",
		"Easter",
		"Ascension Day",
		"Pentecost",
		"", // end sentinel
	}[h]
}

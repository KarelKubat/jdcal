package jdcal

/*
String returns the string representation for a Holiday: "Easter", "Ascension day", etc.
*/
func (h Holiday) String() string {
	return []string{
		"", // start sentinel
		"Easter",
		"Ascension day",
		"Pentecost",
		"", // end sentinel
	}[h]
}

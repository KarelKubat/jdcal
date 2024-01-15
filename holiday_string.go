package jdcal

func (h Holiday) String() string {
	return []string{
		"", // start sentinel
		"Easter",
		"Ascension day",
		"Pentecost",
		"", // end sentinel
	}[h]
}

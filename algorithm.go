package jdcal

var Algorithm ConversionAlgorithm = ByProgression

func (c ConversionAlgorithm) String() string {
	return []string{
		"", // leading sentinel
		"progression",
		"lookup",
		"", // trailing sentinel
	}[c]
}

package jdcal

const (
	algorithmProgression int = iota
	algorithmLookupTable
)

var (
	algorithm int = algorithmProgression
)

/*
ConvertByLookup switches date conversions to a slower, but well tested algorithm which uses a lookup table. This table is derived from  https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars, original source: the Nautical almanac of the United Kingdom and United States (1961).

Conversion by lookup table is NOT the default. If you don't trust the outcome of Date.Convert(), then switch and retry (and let me know if there are problems!).
*/
func ConvertByLookup() {
	algorithm = algorithmLookupTable
}

/*
ConvertByProgression switches date conversions to be faster by using a day progression count since start of epoch (around 500 BC). Converting by progression is the default.
*/
func ConvertByProgression() {
	algorithm = algorithmProgression
}

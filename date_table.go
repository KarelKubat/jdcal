package jdcal

// Generated file, don't edit by hand. Look for a nearby Makefile.

import (
	"time"
)

/*
ConversionTable defines matching Julian and Gregorian dates. It is consulted by, e.g., Convert(). The contained ConversionEntry's won't normally be of use outside of this module, but they are exported (as uppercase symbols) so they can be inspected.

This table reflects https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars, original source: the Nautical almanac of the United Kingdom and United States (1961).

This table
*/
var ConversionTable = [...]ConversionEntry{
	// [...] is syntactic sugar to let the compiler figure out the array size. That way
	// we get a fixed size array and not a slice.
	{
		JDate: Date{Year: -500, Month: time.March, Day: 5, Type: Julian},
		GDate: Date{Year: -500, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: -500, Month: time.March, Day: 6, Type: Julian},
		GDate: Date{Year: -500, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: -5 days
	},
	{
		JDate: Date{Year: -300, Month: time.March, Day: 3, Type: Julian},
		GDate: Date{Year: -300, Month: time.February, Day: 27, Type: Gregorian},
		// Diff: -5 days
	},
	{
		JDate: Date{Year: -300, Month: time.March, Day: 4, Type: Julian},
		GDate: Date{Year: -300, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: -300, Month: time.March, Day: 5, Type: Julian},
		GDate: Date{Year: -300, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: -4 days
	},
	{
		JDate: Date{Year: -200, Month: time.March, Day: 2, Type: Julian},
		GDate: Date{Year: -200, Month: time.February, Day: 27, Type: Gregorian},
		// Diff: -4 days
	},
	{
		JDate: Date{Year: -200, Month: time.March, Day: 3, Type: Julian},
		GDate: Date{Year: -200, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: -200, Month: time.March, Day: 4, Type: Julian},
		GDate: Date{Year: -200, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: -3 days
	},
	{
		JDate: Date{Year: -100, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: -100, Month: time.February, Day: 27, Type: Gregorian},
		// Diff: -3 days
	},
	{
		JDate: Date{Year: -100, Month: time.March, Day: 2, Type: Julian},
		GDate: Date{Year: -100, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: -100, Month: time.March, Day: 3, Type: Julian},
		GDate: Date{Year: -100, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: -2 days
	},
	{
		JDate: Date{Year: 100, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 100, Month: time.February, Day: 27, Type: Gregorian},
		// Diff: -2 days
	},
	{
		JDate: Date{Year: 100, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 100, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: 100, Month: time.March, Day: 2, Type: Julian},
		GDate: Date{Year: 100, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: -1 days
	},
	{
		JDate: Date{Year: 200, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 200, Month: time.February, Day: 27, Type: Gregorian},
		// Diff: -1 days
	},
	{
		JDate: Date{Year: 200, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 200, Month: time.February, Day: 28, Type: Gregorian},
	},
	{
		JDate: Date{Year: 200, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 200, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 0 days
	},
	{
		JDate: Date{Year: 300, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 300, Month: time.February, Day: 28, Type: Gregorian},
		// Diff: 0 days
	},
	{
		JDate: Date{Year: 300, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 300, Month: time.March, Day: 1, Type: Gregorian},
	},
	{
		JDate: Date{Year: 300, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 300, Month: time.March, Day: 2, Type: Gregorian},
		// Diff: 1 days
	},
	{
		JDate: Date{Year: 500, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 500, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 1 days
	},
	{
		JDate: Date{Year: 500, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 500, Month: time.March, Day: 2, Type: Gregorian},
	},
	{
		JDate: Date{Year: 500, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 500, Month: time.March, Day: 3, Type: Gregorian},
		// Diff: 2 days
	},
	{
		JDate: Date{Year: 600, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 600, Month: time.March, Day: 2, Type: Gregorian},
		// Diff: 2 days
	},
	{
		JDate: Date{Year: 600, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 600, Month: time.March, Day: 3, Type: Gregorian},
	},
	{
		JDate: Date{Year: 600, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 600, Month: time.March, Day: 4, Type: Gregorian},
		// Diff: 3 days
	},
	{
		JDate: Date{Year: 700, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 700, Month: time.March, Day: 3, Type: Gregorian},
		// Diff: 3 days
	},
	{
		JDate: Date{Year: 700, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 700, Month: time.March, Day: 4, Type: Gregorian},
	},
	{
		JDate: Date{Year: 700, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 700, Month: time.March, Day: 5, Type: Gregorian},
		// Diff: 4 days
	},
	{
		JDate: Date{Year: 900, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 900, Month: time.March, Day: 4, Type: Gregorian},
		// Diff: 4 days
	},
	{
		JDate: Date{Year: 900, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 900, Month: time.March, Day: 5, Type: Gregorian},
	},
	{
		JDate: Date{Year: 900, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 900, Month: time.March, Day: 6, Type: Gregorian},
		// Diff: 5 days
	},
	{
		JDate: Date{Year: 1000, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1000, Month: time.March, Day: 5, Type: Gregorian},
		// Diff: 5 days
	},
	{
		JDate: Date{Year: 1000, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1000, Month: time.March, Day: 6, Type: Gregorian},
	},
	{
		JDate: Date{Year: 1000, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1000, Month: time.March, Day: 7, Type: Gregorian},
		// Diff: 6 days
	},
	{
		JDate: Date{Year: 1100, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1100, Month: time.March, Day: 6, Type: Gregorian},
		// Diff: 6 days
	},
	{
		JDate: Date{Year: 1100, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1100, Month: time.March, Day: 7, Type: Gregorian},
	},
	{
		JDate: Date{Year: 1100, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1100, Month: time.March, Day: 8, Type: Gregorian},
		// Diff: 7 days
	},
	{
		JDate: Date{Year: 1300, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1300, Month: time.March, Day: 7, Type: Gregorian},
		// Diff: 7 days
	},
	{
		JDate: Date{Year: 1300, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1300, Month: time.March, Day: 8, Type: Gregorian},
	},
	{
		JDate: Date{Year: 1300, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1300, Month: time.March, Day: 9, Type: Gregorian},
		// Diff: 8 days
	},
	{
		JDate: Date{Year: 1400, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1400, Month: time.March, Day: 8, Type: Gregorian},
		// Diff: 8 days
	},
	{
		JDate: Date{Year: 1400, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1400, Month: time.March, Day: 9, Type: Gregorian},
	},
	{
		JDate: Date{Year: 1400, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1400, Month: time.March, Day: 10, Type: Gregorian},
		// Diff: 9 days
	},
	{
		JDate: Date{Year: 1500, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1500, Month: time.March, Day: 9, Type: Gregorian},
		// Diff: 9 days
	},
	{
		JDate: Date{Year: 1500, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1500, Month: time.March, Day: 10, Type: Gregorian},
	},
	{
		JDate: Date{Year: 1500, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1500, Month: time.March, Day: 11, Type: Gregorian},
		// Diff: 10 days
	},
	{
		JDate: Date{Year: 1700, Month: time.February, Day: 18, Type: Julian},
		GDate: Date{Year: 1700, Month: time.February, Day: 28, Type: Gregorian},
		// Diff: 10 days
	},
	{
		JDate: Date{Year: 1700, Month: time.February, Day: 19, Type: Julian},
		GDate: Date{Year: 1700, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 11 days
	},
	{
		JDate: Date{Year: 1700, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1700, Month: time.March, Day: 10, Type: Gregorian},
		// Diff: 11 days
	},
	{
		JDate: Date{Year: 1700, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1700, Month: time.March, Day: 11, Type: Gregorian},
		// Diff: 11 days
	},
	{
		JDate: Date{Year: 1700, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1700, Month: time.March, Day: 12, Type: Gregorian},
		// Diff: 11 days
	},
	{
		JDate: Date{Year: 1800, Month: time.February, Day: 17, Type: Julian},
		GDate: Date{Year: 1800, Month: time.February, Day: 28, Type: Gregorian},
		// Diff: 11 days
	},
	{
		JDate: Date{Year: 1800, Month: time.February, Day: 18, Type: Julian},
		GDate: Date{Year: 1800, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 12 days
	},
	{
		JDate: Date{Year: 1800, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1800, Month: time.March, Day: 11, Type: Gregorian},
		// Diff: 12 days
	},
	{
		JDate: Date{Year: 1800, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1800, Month: time.March, Day: 12, Type: Gregorian},
		// Diff: 12 days
	},
	{
		JDate: Date{Year: 1800, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1800, Month: time.March, Day: 13, Type: Gregorian},
		// Diff: 12 days
	},
	{
		JDate: Date{Year: 1900, Month: time.February, Day: 16, Type: Julian},
		GDate: Date{Year: 1900, Month: time.February, Day: 28, Type: Gregorian},
		// Diff: 12 days
	},
	{
		JDate: Date{Year: 1900, Month: time.February, Day: 17, Type: Julian},
		GDate: Date{Year: 1900, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 13 days
	},
	{
		JDate: Date{Year: 1900, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 1900, Month: time.March, Day: 12, Type: Gregorian},
		// Diff: 13 days
	},
	{
		JDate: Date{Year: 1900, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 1900, Month: time.March, Day: 13, Type: Gregorian},
		// Diff: 13 days
	},
	{
		JDate: Date{Year: 1900, Month: time.March, Day: 1, Type: Julian},
		GDate: Date{Year: 1900, Month: time.March, Day: 14, Type: Gregorian},
		// Diff: 13 days
	},
	{
		JDate: Date{Year: 2100, Month: time.February, Day: 15, Type: Julian},
		GDate: Date{Year: 2100, Month: time.February, Day: 28, Type: Gregorian},
		// Diff: 13 days
	},
	{
		JDate: Date{Year: 2100, Month: time.February, Day: 16, Type: Julian},
		GDate: Date{Year: 2100, Month: time.March, Day: 1, Type: Gregorian},
		// Diff: 14 days
	},
	{
		JDate: Date{Year: 2100, Month: time.February, Day: 28, Type: Julian},
		GDate: Date{Year: 2100, Month: time.March, Day: 13, Type: Gregorian},
		// Diff: 14 days
	},
	{
		JDate: Date{Year: 2100, Month: time.February, Day: 29, Type: Julian},
		GDate: Date{Year: 2100, Month: time.March, Day: 14, Type: Gregorian},
		// Diff: 14 days
	},
}

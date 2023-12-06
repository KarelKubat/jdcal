package jdcal

import "time"

type Type int

const (
	Gregorian Type = iota // jdcal.Gregorian indicates a date on the Gregorian calendar
	Julian                // jdcal.Julian indicates a date on the Julian calendar
)

/*
Date wraps a year, month, day and calendar type (Julian or Gregorian).
*/
type Date struct {
	Year  int
	Month time.Month
	Day   int
	Type  Type
}

/*
ConversionEntry wraps one Julian and one Gregorian date. The ConversionTable is an array of such entries.
*/
type ConversionEntry struct {
	JDate, GDate Date
}

/*
ZoneEntry wraps a zone name with the dates where that zone switched to the Gregorian calendar, possibly later back to the Gregorian, etc. The ZoneTable is an array of such entries.

The CutOvers list is an array of Dates where this zone switched FROM a calendar TO another one. E.g., when CutOvers is:

	Cutovers: []Date{
		{Year: -500, Month: time.February, Day: 28, Type: Gregorian},
		{Year: 1584, Month: time.January, Day: 1, Type: Julian},
		{Year: 1597, Month: time.January, Day: 1, Type: Gregorian},
		{Year: 1798, Month: time.December, Day: 25, Type: Julian},
	}

Then that means:

- The calendar for this zone starts at -500/02/28 Gregorian (or: -500 March 1st Julian, convert yourself if you think that is handy). On that day the zone started using the Julian calendar. This seems counter-intuitive, but "started using" is just by definition the "other calendar than this entry".

- On 1584/01/01 Julian, the zone switched to Gregorian ("the other one").

- On 1597/01/01 Gregorian, the zone switched to Julian. So they switched back.

- Finally, on 1798/12/25, they switched again. To Gregorian.
*/
type ZoneEntry struct {
	Name     string // Zone name, e.g. "Denmark"
	Cutovers []Date // List of dates when the zone switched from a given calendar
}

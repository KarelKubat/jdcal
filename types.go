package jdcal

import "time"

/*
Type (an int) defines the calendar type for a jdcal.Date or a jdcal.CalendarYear:
jdcal.Gregorian or jdcal.Julian.
*/
type Type int

const (
	Gregorian Type = iota // jdcal.Gregorian indicates a date on the Gregorian calendar
	Julian                // jdcal.Julian indicates a date on the Julian calendar
)

/*
MD is a representation of a month and a day; basically a YMD without a year.
*/
type MD struct {
	Month time.Month
	Day   int
}

/*
YMD is a representation of a year, month and a day; basically a Date without a type.
*/
type YMD struct {
	Year  Year
	Month time.Month
	Day   int
}

/*
Date wraps a year, month, day and calendar type (Julian or Gregorian); basically a YMD with a type.
*/
type Date struct {
	Year  Year
	Month time.Month
	Day   int
	Type  Type
}

/*
Year is an integer and the receiver type for some helper functions.
*/
type Year int

/*
CalendarYear wraps a Year with a calendar Type (Julian or Gregorian).
*/
type CalendarYear struct {
	Year Year
	Type Type
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
		// -500 to 1584: Julian calendar applies
		// 1584 to 1597: Switched to Gregorian
		// 1597 to 1798: Switched back to Julian
		// 1798 to now:  Ended up with Gregorian
		{Year: -500, Month: time.February, Day: 28, Type: Gregorian},
		{Year: 1584, Month: time.January, Day: 1, Type: Julian},
		{Year: 1597, Month: time.January, Day: 1, Type: Gregorian},
		{Year: 1798, Month: time.December, Day: 25, Type: Julian},
	}

Then that means:

- The calendar for this zone starts at -500/02/28 Gregorian (or: -500 March 1st Julian, convert yourself if you think that is handy). On that day the zone stopped using the Gregorian, and started using the Julian calendar.

- On 1584/01/01 Julian, the zone switched to Gregorian ("the other one").

- On 1597/01/01 Gregorian, the zone switched to Julian. So they switched back.

- Finally, on 1798/12/25, they switched again. To Gregorian.
*/
type ZoneEntry struct {
	Name     string // Zone name, e.g. "Denmark"
	Cutovers []Date // List of dates when the zone switched from a given calendar
}

/*
Holiday enumarates yearly holidays.
*/
type Holiday int

const (
	firstUnusedHoliday Holiday = iota
	AshWednesday
	GoodFriday
	Easter
	Ascension
	Pentecost
	lastUnusedHoliday
)

/*
Ordinal (an int) represents the day number since epoch start.
*/
type Ordinal int

/*
ConversionAlgoritm (an int) represents the applicable conversion algorithm.
*/
type ConversionAlgorithm int

const (
	firstUnusedAlgorithm ConversionAlgorithm = iota
	ByProgression                            // Default conversion algorithm, fast
	ByLookup                                 // Slow conversion algorithm, but thoroughly tested
	lastUnusedAlgorithm
)

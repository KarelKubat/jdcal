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

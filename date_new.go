package jdcal

import (
	"time"
)

/*
NewDate is a helper function to construct a Date from a year, month, day and calendar type.
The two snippets are equivalent:

	d, err := jdcal.New(1962, time.August, 19, jdcal.Gregorian)
	if err != nil {...}

	d := Date{Year: 1962, Month: time.August, Day: 19, Type: jdcal.Gregorian}
	err := d.Valid()
	if err {...}
*/
func NewDate(year Year, month time.Month, day int, tp Type) (dt Date, err error) {
	dt = Date{
		Year:  year,
		Month: month,
		Day:   day,
		Type:  tp,
	}
	return dt, dt.Valid()
}

/*
NewFromString is a helper function to convert a string in the format "YYYY/MM/DD" into a date of a given Type (Julian or Gregorian). It chaines StringToYMD() and New().
*/
func NewDateFromString(arg string, tp Type) (dt Date, err error) {
	ymd, err := StringToYMD(arg)
	if err != nil {
		return dt, err
	}
	return NewDate(ymd.Year, ymd.Month, ymd.Day, tp)
}

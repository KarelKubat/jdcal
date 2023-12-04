package jdcal

import "time"

/*
New is a helper function to construct a Date from a year, month, day and calendar type.
The two snippets are equivalent:

  d, err := jdcal.New(1962, time.August, 19, jdcal.Gregorian)
  if err != nil {...}

  d := Date{Year: 1962, Month: time.August, Day: 19, Type: jdcal.Gregorian}
  err := d.Valid()
  if err {...}
*/
func New(year int, month time.Month, day int, tp Type) (Date, error) {
	dt := Date{
		Year:  year,
		Month: month,
		Day:   day,
		Type:  tp,
	}
	return dt, dt.Valid()
}

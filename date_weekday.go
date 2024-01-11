package jdcal

import "time"

/*
Weekday returns the day of the week (time.Sunday, time.Monday etc.) for the given date.
The date type (jdcal.Julian or jdcal.Gregorian) doesn't matter; if needed, the date is converted
before determining the weekday.

The following two snippets are equivalent:

	dt, err := jdcal.New(...)
	if err != nil { ... }

	// Alternative 1
	wd, err := dt.Weekday()
	if err != nil { ... }

	// Alternative 2
	tm, err := dt.AsTime()
	if err != nil { ... }
	wd := tm.Weekday()
*/
func (d Date) Weekday() (wd time.Weekday, err error) {
	tm, err := d.AsTime()
	if err != nil {
		return time.Sunday, err
	}
	return tm.Weekday(), nil
}

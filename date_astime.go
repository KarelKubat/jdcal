package jdcal

import "time"

/*
AsTime returns the jdcal.Date as a time.Time type, so that one can apply all standard functions of
https://pkg.go.dev/time#Time. The time coordinates are pinned at noon, UTC.

The date type (jdcal.Julian or jdcal.Gregorian) does not matter; the date is converted if needed
(time.Time is, by definition, Gregorian).
*/
func (d Date) AsTime() (tm time.Time, err error) {
	var ot Date
	if d.Type == Julian {
		ot, err = d.Convert()
		if err != nil {
			return time.Time{}, err
		}
	} else {
		ot = d
	}
	return time.Date(ot.Year, ot.Month, ot.Day, 12, 0, 0, 0, time.UTC), nil
}

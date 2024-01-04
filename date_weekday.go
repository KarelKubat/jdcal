package jdcal

import "time"

func (d Date) Weekday() (wd time.Weekday, err error) {
	var ot Date
	if d.Type == Julian {
		ot, err = d.Convert()
		if err != nil {
			return time.Sunday, err
		}
	} else {
		ot = d
	}
	return time.Date(ot.Year, ot.Month, ot.Day, 12, 0, 0, 0, time.UTC).Weekday(), nil
}

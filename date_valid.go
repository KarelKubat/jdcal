package jdcal

import (
	"fmt"
	"time"
)

/*
Valid returns an error when a date cannot be processed. The date must not exceed the maximum number of month days (e.g., April 31st is wrong, February 29th may only occur in leap years) and Convert() must be able to process it: it can't be outside of the range [Date.First() .. Date.Last()].
*/
func (d Date) Valid() error {
	// Year verification
	if err := d.Year.Valid(); err != nil {
		return err
	}

	// Day verification, [1..limit-per-month]
	if d.Day < 1 {
		return fmt.Errorf("%v: day must be >=1", d)
	}
	maxPerMonth := []int{
		0, // Filler as month.January is 1
		31, 29, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	if d.Day > maxPerMonth[d.Month] {
		return fmt.Errorf("%v can't have %v days", d.Month, d.Day)
	}

	// Even with the max per month, Feb 28'th may not be valid.
	if d.Month == time.February && d.Day == 29 {
		cyr, err := NewCalendarYear(d.Year, d.Type)
		if err != nil {
			panic(fmt.Sprintf("internal error: Date.Valid failed to construct a CalendarYear: %v", err))
		}
		isLeap := cyr.IsLeap()
		if !isLeap {
			return fmt.Errorf("year %v has no February 29th", d.Year)
		}
	}

	// Dates can't exceed the conversion table.
	before, err := d.Before(First(d.Type))
	if err != nil {
		return err
	}
	if before {
		return fmt.Errorf("%v %s (%v)", d, beforeConvertibleDate, First(d.Type))
	}
	after, err := d.After(Last(d.Type))
	if err != nil {
		return err
	}
	if after {
		return fmt.Errorf("%v %s (%v)", d, afterConvertibleDate, (d.Type))
	}

	return nil
}

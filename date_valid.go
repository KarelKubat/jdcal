package jdcal

import (
	"fmt"
	"time"
)

/*
Valid returns an error when a date cannot be processed. The date must not exceed the maximum number of month days (e.g., April 31st is wrong, February 29th may only occur in leap years) and Convert() must be able to process it (it can't be outside of the ConversionTable).
*/
func (d Date) Valid() error {
	// Year verification
	if err := d.Year.Valid(); err != nil {
		return err
	}

	// There is a limit to the # of month days (e.g. Feb can't have 30 days).
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
		isLeap := CalendarYear{Year: d.Year, Type: d.Type}.IsLeap()
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

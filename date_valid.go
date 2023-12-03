package jdcal

import (
	"fmt"
	"time"
)

func (d Date) Valid() error {
	// There is a limit to the # of month days (e.g. Feb can't have 30 days).
	maxPerMonth := []int{
		31, 29, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	if d.Day > maxPerMonth[d.Month] {
		return fmt.Errorf("%v can't have %v days", d.Month, d.Day)
	}

	// Even with the max per month, Feb 28'th may not be valid.
	if d.Month == time.February && d.Day == 29 && !d.IsLeap() {
		return fmt.Errorf("year %4.4d has no February 29th", d.Year)
	}

	// Dates can't exceed the conversion table.
	before, err := d.Before(First(d.Type))
	if err != nil {
		return err
	}
	if before {
		return fmt.Errorf("%v is before the first convertible date (%v)",
			d, First(d.Type))
	}
	after, err := d.After(Last(d.Type))
	if err != nil {
		return err
	}
	if after {
		return fmt.Errorf("%v is after the last convertible date (%v)", d, Last(d.Type))
	}

	return nil
}

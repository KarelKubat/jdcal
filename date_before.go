package jdcal

import "errors"

/*
Before is true when the date in question occurs earlier than the other date. Note that
different date types raise an error, see After().
*/
func (d Date) Before(other Date) (bool, error) {
	if d.Type != other.Type {
		return false, errors.New("can't compare different types")
	}

	if d.Year < other.Year {
		return true, nil
	}
	if d.Year > other.Year {
		return false, nil
	}

	if d.Month < other.Month {
		return true, nil
	}
	if d.Month > other.Month {
		return false, nil
	}

	if d.Day < other.Day {
		return true, nil
	}
	return false, nil
}

package jdcal

import (
	"errors"
)

func (d Date) After(other Date) (bool, error) {
	if d.Type != other.Type {
		return false, errors.New("can't compare different types")
	}

	if d.Year > other.Year {
		return true, nil
	}
	if d.Year < other.Year {
		return false, nil
	}

	if d.Month > other.Month {
		return true, nil
	}
	if d.Month < other.Month {
		return false, nil
	}

	if d.Day > other.Day {
		return true, nil
	}
	return false, nil
}

package jdcal

import "errors"

func (d Date) Equal(other Date) (bool, error) {
	if d.Type != other.Type {
		return false, errors.New("can't compare different types")
	}

	return d.Year == other.Year && d.Month == other.Month && d.Day == other.Day, nil
}

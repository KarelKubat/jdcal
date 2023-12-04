package jdcal

import "errors"

/*
Equal is true when two dates point to the same day. Different date types cannot be compared,
see After().
*/
func (d Date) Equal(other Date) (bool, error) {
	if d.Type != other.Type {
		return false, errors.New("can't compare different types")
	}

	return d.Year == other.Year && d.Month == other.Month && d.Day == other.Day, nil
}

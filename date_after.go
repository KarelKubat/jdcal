package jdcal

import (
	"errors"
)

/*
After is true when the date in question occurs later than the other date. Note that different
date types raise an error. When comparing different date types, the caller must first convert:

	jd, err := jdcal.New(1666, time.March, 13, jdcal.Julian) // a Julian date in 1666
	if err != nil {...}

	gd, err := jdcal.New(2023, time.December, 4, jdcal.Gregorian) // a Gregorian date in 2023
	if err != nil {...}

	gdTmp, err := jd.Convert() // jd as a Gregorian date
	if err != nil {...}
	fmt.Println(gd.After(gdTmp)) // true; the 2023 date comes after the 1666 date
*/
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

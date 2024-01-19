package jdcal

import (
	"errors"
	"fmt"
)

const (
	// Error texts shared with date_valid.go
	beforeConvertibleDate = "is before the first convertible date"
	afterConvertibleDate  = "is after the last convertible date"
)

/*
Valid returns an error when a year points outside of the known dates of the conversion
table, or when the year is zero (there is no zero in BC/AD).
*/
func (y Year) Valid() error {
	if y == 0 {
		return errors.New("year 0 does not exist in BC/AD")
	}
	switch algorithm {
	case algorithmLookupTable:
		if y < First(Julian).Year {
			return fmt.Errorf("year %v %s", y, beforeConvertibleDate)
		}
		if y > Last(Gregorian).Year {
			return fmt.Errorf("year %v %s", y, afterConvertibleDate)
		}
		return nil
	case algorithmProgression:
		if y < StartProgressionYear {
			return fmt.Errorf("year %v %s", y, beforeConvertibleDate)
		}
		if y > EndProgressionYear {
			return fmt.Errorf("year %v %s", y, afterConvertibleDate)
		}
		return nil
	default:
		return errors.New("algorithm mismatch in Year.Valid")
	}
}

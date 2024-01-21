package jdcal

import (
	"errors"
	"fmt"
)

/*
NewCalendarYear returns a CalendarYear for a given year and calendar type.
*/
func NewCalendarYear(y Year, tp Type) (CalendarYear, error) {
	if y == 0 {
		return CalendarYear{}, errors.New("year zero does not exist in BC/AD")
	}
	if y < First(tp).Year || y > Last(tp).Year {
		return CalendarYear{}, fmt.Errorf("year %d out of range", y)
	}
	return CalendarYear{Year: y, Type: tp}, nil
}

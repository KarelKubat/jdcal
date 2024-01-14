package jdcal

import "fmt"

/*
String returns the readable representation for a CalendarYear, e.g. "1600 Julian".
*/

func (c CalendarYear) String() string {
	return fmt.Sprintf("%v %v", c.Type, c.Year)
}

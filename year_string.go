package jdcal

import "fmt"

/*
String returns the string representation for a year, left-padded with zeroes over four
positions, and if needed prefixed with a minus sign.
*/
func (y Year) String() string {
	return fmt.Sprintf("%4.4d", y)
}

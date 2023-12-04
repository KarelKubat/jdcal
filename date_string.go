package jdcal

import "fmt"

/*
String returns a printable version of a date in the format "YYYY/MM/DD TYPE", e.g. "1234/12/27 Julian".  The year can be prefixed by a - to indicate negative values.
*/
func (d Date) String() string {
	return fmt.Sprintf("%s %4.4d/%2.2d/%2.2d", d.Type, d.Year, d.Month, d.Day)
}

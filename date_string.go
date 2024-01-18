package jdcal

import "fmt"

/*
String returns a printable version of a date in the format "TYPE YYYY/MM/DD", e.g. "Julian 1234/12/27".
The year may be prefixed by a - to indicate negative values.
*/
func (d Date) String() string {
	return fmt.Sprintf("%s %s", d.Type, YMD{d.Year, d.Month, d.Day})
}

package jdcal

import "fmt"

/*
String returns a string representation of an MD in the form MM/DD.
*/

func (m MD) String() string {
	return fmt.Sprintf("%2.2d/%2.2d", m.Month, m.Day)
}

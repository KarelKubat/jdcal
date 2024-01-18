package jdcal

import "fmt"

func (y YMD) String() string {
	return fmt.Sprintf("%v/%2.2d/%2.2d", y.Year, y.Month, y.Day)
}

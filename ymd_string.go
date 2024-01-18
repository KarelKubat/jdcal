package jdcal

import "fmt"

func (y YMD) String() string {
	return fmt.Sprintf("%v/%v", y.Year, MD{Month: y.Month, Day: y.Day})
}

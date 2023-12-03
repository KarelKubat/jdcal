package jdcal

import "fmt"

func (d *Date) String() string {
	return fmt.Sprintf("%s %4.4d/%2.2d/%2.2d", d.Type, d.Year, d.Month, d.Day)
}

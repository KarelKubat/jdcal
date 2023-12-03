package jdcal

import (
	"time"
)

func (d *Date) Advance() {
	// Nr of month days, with 28 for February (handled separately).
	daysPerMonth := []int{
		0, // filler, January is #1
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}

	if d.Day < daysPerMonth[d.Month] {
		d.Day++
		return
	}
	if d.Month == time.February {
		if d.Day == 28 && d.IsLeap() {
			d.Day++
		} else {
			d.Month = time.March
			d.Day = 1
		}
		return
	}
	if d.Month == time.December {
		d.Year++
		d.Day = 1
		d.Month = time.January
		return
	}
	d.Day = 1
	d.Month++
}

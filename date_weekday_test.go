package jdcal

import (
	"time"

	"testing"
)

func TestWeekday(t *testing.T) {
	for _, test := range []struct {
		day         int
		wantWeekday time.Weekday
	}{
		{
			day:         1,
			wantWeekday: time.Monday,
		},
		{
			day:         2,
			wantWeekday: time.Tuesday,
		},
		{
			day:         3,
			wantWeekday: time.Wednesday,
		},
		{
			day:         4,
			wantWeekday: time.Thursday,
		},
		{
			day:         5,
			wantWeekday: time.Friday,
		},
		{
			day:         6,
			wantWeekday: time.Saturday,
		},
		{
			day:         7,
			wantWeekday: time.Sunday,
		},
		{
			day:         8,
			wantWeekday: time.Monday,
		},
	} {
		d, err := New(2024, time.January, test.day, Gregorian)
		if err != nil {
			t.Fatalf("New(...) = _,%q, need nil error", err.Error())
		}
		wd, err := d.Weekday()
		if err != nil {
			t.Fatalf("Weekday() = _,%q, need nil error", err.Error())
		}
		if wd != test.wantWeekday {
			t.Errorf("%+v vWeekday() = %v, want %v", d, wd, test.wantWeekday)
		}
	}
}

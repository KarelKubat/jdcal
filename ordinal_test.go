package jdcal

import (
	"testing"
	"time"
)

func TestOrdinal(t *testing.T) {
	// See also main/demo7/demo7.go
	for _, ymd := range []YMD{
		// Near start of epoch
		{Year: -500, Month: time.January, Day: 1},

		// Around negative leap
		{Year: -301, Month: time.February, Day: 28},
		{Year: -301, Month: time.March, Day: 1},

		// Aroound positive leap
		{Year: 300, Month: time.February, Day: 28},
		{Year: 300, Month: time.March, Day: 1},

		// Somewhere in the 20th century
		{Year: 1962, Month: time.August, Day: 19},

		// Near end of epoch
		{Year: 2100, Month: time.January, Day: 1},
	} {
		for _, tp := range []Type{Gregorian, Julian} {
			dt, err := New(ymd.Year, ymd.Month, ymd.Day, tp)
			check(t, err)
			ord, err := dt.Ordinal()
			check(t, err)
			back, err := ord.Date(dt.Type)
			check(t, err)

			eq, err := dt.Equal(back)
			check(t, err)
			if !eq {
				t.Errorf("%v -> %6d -> %v: mismatch", dt, ord, back)
			}

		}
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

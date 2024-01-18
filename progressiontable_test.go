package jdcal

import (
	"testing"
	"time"
)

func TestMonthProgressionTables(t *testing.T) {
	for _, monthProgression := range []map[time.Month][]int{
		LeapMonthProgression,
		NonLeapMonthProgression,
	} {
		for m := time.January; m <= time.December; m++ {
			if monthProgression[m][0] != 0 {
				t.Fatalf("missing zero leader in table %v", monthProgression)
			}
		}
	}

	febDays := LeapMonthProgression[time.February]
	if febDays[len(febDays)-1] != 60 { // jan: 31 plus feb: 29
		t.Error("LeapMonthProgression: last february entry in leap year must be 29")
	}
	febDays = NonLeapMonthProgression[time.February]
	if febDays[len(febDays)-1] != 59 { // jan: 31, feb: 28
		t.Error("NonLeapMonthProgression: last february entry in nonleap year must be 59")
	}
	decDays := LeapMonthProgression[time.December]
	if decDays[len(decDays)-1] != 366 { // leap year: 366 days
		t.Error("LeapMonthProgression: last december entry in leap year must be 366")
	}
	decDays = NonLeapMonthProgression[time.December]
	if decDays[len(decDays)-1] != 365 { // std year: 365 days
		t.Error("NonLeapMonthProgression: last december entry in nonleap year must be 365")
	}
}

func TestYearProgressionTable(t *testing.T) {
	for _, test := range []struct {
		year          Year
		wantGregorian int
		wantJulian    int
	}{
		// Bunch of canned progressions.

		// Table start
		{
			year:          -500,
			wantGregorian: 0,
			wantJulian:    -5,
		},

		// Negative leap year difference
		{
			year:          -101,
			wantGregorian: 145732,
			wantJulian:    145729,
		},
		{
			year:          -100,
			wantGregorian: 146097,
			wantJulian:    146095,
		},

		// Converging calendars
		{
			year:          219,
			wantGregorian: 262610,
			wantJulian:    262610,
		},

		// Positive leap year difference
		{
			year:          1700,
			wantGregorian: 803535,
			wantJulian:    803545,
		},
	} {
		if got := YearProgression[test.year][Gregorian]; got != test.wantGregorian {
			t.Errorf("YearProgression[%d][Gregorian] = %v, want %v", test.year, got, test.wantGregorian)
		}
		if got := YearProgression[test.year][Julian]; got != test.wantJulian {
			t.Errorf("YearProgression[%d][Julian] = %v, want %v", test.year, got, test.wantJulian)
		}
	}
}

package jdcal

import (
	"testing"
	"time"
)

func TestMonthProgressionTables(t *testing.T) {
	for _, monthProgression := range []MonthProgression{
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
	if febDays[len(febDays)-1] != 59 { // jan: 31 plus feb: 29 = 60, minus 1
		t.Error("LeapMonthProgression: last february entry in leap year must be 59")
	}
	febDays = NonLeapMonthProgression[time.February]
	if febDays[len(febDays)-1] != 58 { // jan: 31, feb: 28 = 59, minus 1
		t.Error("NonLeapMonthProgression: last february entry in nonleap year must be 58")
	}
	decDays := LeapMonthProgression[time.December]
	if decDays[len(decDays)-1] != 365 { // leap year: 366 days, minus 1
		t.Error("LeapMonthProgression: last december entry in leap year must be 366")
	}
	decDays = NonLeapMonthProgression[time.December]
	if decDays[len(decDays)-1] != 364 { // std year: 365 days, minus 1
		t.Error("NonLeapMonthProgression: last december entry in nonleap year must be 364")
	}
}

func TestYearProgressionTable(t *testing.T) {
	for _, test := range []struct {
		year          Year
		wantGregorian Ordinal
		wantJulian    Ordinal
	}{
		// Bunch of canned progressions. These tests break when the table is generated with different
		// start/end years by main/makeprogressiontable/makeprogressiontable.sh.

		// Near table start
		{
			year:          -500,
			wantGregorian: 2191,
			wantJulian:    2186,
		},

		// Near table end
		{
			year:          2100,
			wantGregorian: 951823,
			wantJulian:    951836,
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

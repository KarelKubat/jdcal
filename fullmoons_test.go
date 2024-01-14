package jdcal

import "testing"

func TestFullMoonsTable(t *testing.T) {
	nYears := 0

	for yr, dates := range FullMoons {
		nYears++

		// Catch empty entries
		if len(dates) < 2 {
			t.Errorf("FullMoons table: year %d: not enough dates", yr)
		}
	}

	// Catch bad generation
	if nYears < 2500 {
		t.Errorf("FullMoons table: found only %d years", nYears)
	}
}

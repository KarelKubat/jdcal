package jdcal

import (
	"testing"
	"time"
)

func TestFullMoonsTable(t *testing.T) {
	nYears := 0

	for yr, date := range FullMoons {
		nYears++

		// Month / day must be beyond March 21st.
		if date.Month < time.March || (date.Month == time.March && date.Day < 21) {
			// Catch empty entries
			t.Errorf("FullMoons table: year %d: date %v is before March 21st", yr, date)
		}
	}

	// Catch bad generation
	if nYears < 2500 {
		t.Errorf("FullMoons table: found only %d years", nYears)
	}
}

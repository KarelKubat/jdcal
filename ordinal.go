package jdcal

import (
	"fmt"
	"time"
)

/*
OrdinalToDate returns the jdcal.Date for a given ordinal day number. This is the reverse of
Date.Ordinal.
*/
func OrdinalToDate(ordinal int, tp Type) (Date, error) {
	var yr Year
	for yr = StartProgressionYear; yr <= EndProgressionYear; yr++ {
		if yr == EndProgressionYear {
			break
		}
		nextProgression := YearProgression[yr+1][tp]
		if nextProgression > ordinal {
			break
		}
	}

	var progression map[time.Month][]int
	cyr := CalendarYear{Year: yr, Type: tp}
	if cyr.IsLeap() {
		progression = LeapMonthProgression
	} else {
		progression = NonLeapMonthProgression
	}

	// Start looking at the base number of the overall years progression.
	ord := YearProgression[yr][tp]

	// Find the month that reduces ordinal minus ord the best.
	var m time.Month
	for m = time.January; m <= time.December; m++ {
		max := lastInt(progression[m])
		if ordinal-ord-max <= 0 {
			break
		}
	}

	// Find the day to fill the gap.
	for nr, d := range progression[m] {
		if ord+d >= ordinal {
			return New(yr, m, nr, tp)
		}
	}
	return Date{}, fmt.Errorf("failed to find date for ordinal %d, type %s", ordinal, tp)
}

// Helper to return the last int of a slice.
func lastInt(sl []int) int {
	if len(sl) == 0 {
		// Fail fast, fail hard
		panic("FIXME! lastInt called with zero length int slice")
	}
	return sl[len(sl)-1]
}

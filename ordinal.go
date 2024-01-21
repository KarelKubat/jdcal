package jdcal

import (
	"fmt"
	"time"
)

/*
Date returns the jdcal.Date for a given ordinal day number. This is the reverse of
Date.Ordinal.
*/
func (o Ordinal) Date(tp Type) (Date, error) {
	// Find the year where the ordinal equals the requested one, or where it's (just) below.
	yr := o.Year(tp)
	// fmt.Println("OrdinalToDate, ordinal:", o, "for type", tp, "gives yr:", yr)

	// Ordinal days for that year are either by nonleap or by leap.
	var progression MonthProgression
	cyr, err := NewCalendarYear(yr, tp)
	if err != nil {
		return Date{}, err
	}
	if cyr.IsLeap() {
		progression = LeapMonthProgression
	} else {
		progression = NonLeapMonthProgression
	}

	// Start looking at the base number of the overall years progression.
	ord := YearProgression[yr][tp]

	// First/last edge cases
	missing := o - ord
	if missing == 0 {
		return NewDate(yr, time.January, 1, tp)
	}
	var max Ordinal
	if cyr.IsLeap() {
		max = 365
	} else {
		max = 364
	}
	if missing == max {
		return NewDate(yr, time.December, 31, tp)
	}

	// Find the month that reduces ordinal minus ord the best.
	var m time.Month
	for m = time.January; m <= time.December; m++ {
		max := lastOrd(progression[m])
		if missing-max <= 0 {
			break
		}
	}
	// fmt.Println("OrdinalToDate for", o, ": year start", ord, "missing days:", missing, "best month:", m)

	// Find the day to fill the gap.
	for day, count := range progression[m] {
		if day == 0 {
			continue // Skip fillers
		}
		// fmt.Println("OrdinalToDate: day", day, "count", count, "missing", missing)
		if count >= missing {
			return NewDate(yr, m, day, tp)
		}
	}
	return Date{}, fmt.Errorf("failed to find date for ordinal %d, type %s", o, tp)
}

// Helper to return the last int of a slice.
func lastOrd(sl []Ordinal) Ordinal {
	if len(sl) == 0 {
		// Fail fast, fail hard
		panic("internal error: lastOrd called with zero length int slice")
	}
	return sl[len(sl)-1]
}

/*
Year returns the best matching year for a given ordinal.
*/
func (o Ordinal) Year(tp Type) Year {
	for yr := progressionTableEnd; yr >= progressionTableStart; yr-- {
		if YearProgression[Year(yr)][tp] <= o {
			return Year(yr)
		}
	}
	panic(fmt.Sprintf("internal error: Ordinal.Year failed to find a year for ordinal %d, type %v", o, tp))
}

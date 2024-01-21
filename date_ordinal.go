package jdcal

import "fmt"

/*
Ordinal returns the ordinal day for a given date. This is a daycount since "start of epoch", liberally defined as the constant StartProgressionYear. This is the reverse of OrdinalToDate.
*/
func (d Date) Ordinal() Ordinal {
	var progression MonthProgression

	cyr, err := NewCalendarYear(d.Year, d.Type)
	if err != nil {
		panic(fmt.Sprintf("internal error: Date.Ordinal failed to construct a CalendarYear: %v", err))
	}

	if cyr.IsLeap() {
		progression = LeapMonthProgression
	} else {
		progression = NonLeapMonthProgression
	}

	ordinal := YearProgression[d.Year][d.Type] + progression[d.Month][d.Day]
	return ordinal
}

package jdcal

/*
Ordinal returns the ordinal day for a given date. This is a daycount since "start of epoch", liberally defined as the constant StartProgressionYear. This is the reverse of OrdinalToDate.
*/
func (d Date) Ordinal() (Ordinal, error) {
	var progression MonthProgression

	cyr := CalendarYear{Year: d.Year, Type: d.Type}
	if err := cyr.Year.Valid(); err != nil {
		return 0, err
	}

	if cyr.IsLeap() {
		progression = LeapMonthProgression
	} else {
		progression = NonLeapMonthProgression
	}

	ordinal := YearProgression[d.Year][d.Type] + progression[d.Month][d.Day]
	return ordinal, nil
}

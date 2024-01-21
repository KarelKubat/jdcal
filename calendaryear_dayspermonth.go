package jdcal

/*
DaysPerMonth returns for each time.Month the number of days, taking into account the calendar type (Julian or Gregorian) and leap years. Example:

	cyr, err := NewCalendarYear(1900, jdcal.Julian)
	if err != nil { ... }
	dpm := cyr..DaysPerMonth()
	fmt.Println(dpm[time.February])
*/
func (c CalendarYear) DaysPerMonth() []int {
	daysPerMonth := []int{
		0, // filler, January is #1
		31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31,
	}
	if c.IsLeap() {
		daysPerMonth[2] = 29
	}
	return daysPerMonth
}

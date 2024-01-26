package jdcal

import (
	"fmt"
	"time"
)

/*
HolidayDate returns for a given CalendarYear the Date of one of the Holidays. The type of the returned date matches the type of the calendar year. E.g., if the calendar year is Julian 1370, then the returned date is also on the Julian calendar.

Easter is defined as the next Sunday beyond the ecclesiastical spring equinox of March 21st. If this equinox is a Sunday, then the next one is taken.

Ascension Day is on the 40th day after Easter (so plus 39), making it a Thursday.

Pentecost is on the 50th day after Easter (so plus 49), making it a Sunday.
*/
func (cyr CalendarYear) HolidayDate(h Holiday) (Date, error) {
	switch h {
	case AshWednesday:
		return cyr.ashwednesday()
	case GoodFriday:
		return cyr.goodfriday()
	case Easter:
		return cyr.easter()
	case Ascension:
		return cyr.ascension()
	case Pentecost:
		return cyr.pentecost()
	}

	return Date{}, fmt.Errorf("unmached holiday type %v in HolidayDate", h)
}

func (cyr CalendarYear) easter() (Date, error) {
	// Find the first full moon on or beyond the ecclesiastical spring equinox (March 21st).
	fullMoon, ok := FullMoons[cyr.Year]
	if !ok {
		return Date{}, fmt.Errorf("no spring full moon information for %v", cyr)
	}
	dt := Date{Year: cyr.Year, Month: fullMoon.Month, Day: fullMoon.Day, Type: Gregorian}

	// Advance until the next Sunday, but don't take this date if it's a Sunday itself.
	dt = dt.Forward()
	for {
		wd, err := dt.Weekday()
		if err != nil {
			return Date{}, err
		}
		if wd == time.Sunday {
			break
		}
		dt = dt.Forward()
	}
	// fmt.Println("sunday after:", dt)

	// Return the found date, ensuring the requested type.
	var err error
	if dt.Type != cyr.Type {
		dt, err = dt.Convert()
		if err != nil {
			return Date{}, err
		}
	}
	return dt, nil
}

func (cyr CalendarYear) ashwednesday() (Date, error) {
	dt, err := cyr.easter()
	if err != nil {
		return Date{}, err
	}

	// Ash Wednesday is 46 days back.
	for i := 0; i < 46; i++ {
		dt = dt.Backward()
	}
	return dt, nil

}

func (cyr CalendarYear) goodfriday() (Date, error) {
	dt, err := cyr.easter()
	if err != nil {
		return Date{}, err
	}

	// Good Friday is 2 days back.
	for i := 0; i < 2; i++ {
		dt = dt.Backward()
	}
	return dt, nil
}

func (cyr CalendarYear) ascension() (Date, error) {
	dt, err := cyr.easter()
	if err != nil {
		return Date{}, err
	}

	// Ascension is on the 40th day after Easter, so plus 39, a Thursday.
	for i := 0; i < 39; i++ {
		dt = dt.Forward()
	}
	return dt, nil
}

func (cyr CalendarYear) pentecost() (Date, error) {
	dt, err := cyr.easter()
	if err != nil {
		return Date{}, err
	}

	// Pentecost is on the 50th day after Easter, so plus 49, a Sunday.
	for i := 0; i < 49; i++ {
		dt = dt.Forward()
	}
	return dt, nil
}

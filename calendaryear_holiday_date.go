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
	fullMoons, ok := FullMoons[cyr.Year]
	if !ok {
		return Date{}, fmt.Errorf("no full moons information for %v", cyr)
	}

	// Equinox date is on March 21st, either Julian or Gregorian.
	// If the date is requested as Julian, then the search date into the full moons table must
	// be set as Gregorian.
	var err error
	equinox := Date{Year: cyr.Year, Month: time.March, Day: 21, Type: cyr.Type}
	if cyr.Type == Julian {
		equinox, err = equinox.Convert()
		if err != nil {
			return Date{}, err
		}
	}
	// fmt.Println("equinox pinned to:", equinox)
	var dt Date
	var found bool
	for _, dt = range fullMoons {
		// March 21st exactly, such as 1370? Or just beyond, such as April 8th 1371?
		found, err = dt.AfterOrEqual(equinox)
		if err != nil {
			return Date{}, err
		}
		if found {
			break
		}
	}
	if !found {
		return Date{}, fmt.Errorf("failed to find the fullmoon beyond March 21st in year %v", cyr.Year)
	}
	// fmt.Println("fullmoon after equinox:", dt)

	// Found the full moon date. Advance until the next Sunday, but don't take this date if it's
	// a Sunday itself.
	dt = dt.Advance()
	for {
		wd, err := dt.Weekday()
		if err != nil {
			return Date{}, err
		}
		if wd == time.Sunday {
			break
		}
		dt = dt.Advance()
	}
	// fmt.Println("sunday after:", dt)

	// Return the found date, ensuring the requested type.
	if dt.Type != cyr.Type {
		dt, err = dt.Convert()
		if err != nil {
			return Date{}, err
		}
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
		dt = dt.Advance()
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
		dt = dt.Advance()
	}
	return dt, nil
}

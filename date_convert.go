package jdcal

import (
	"errors"
)

func (d Date) Convert() (Date, error) {
	for i, e := range ConversionTable {
		var our, their, ourNext Date
		if d.Type == Gregorian {
			our = e.GDate
			their = e.JDate
		} else {
			our = e.JDate
			their = e.GDate
		}

		// Did we hit a lookup?
		eq, err := d.Equal(our)
		if err != nil {
			return d, err
		}
		if eq {
			return their, nil
		}

		// Can't arrive at the last entry, that's a lookup.
		if i == len(ConversionTable)-1 {
			return d, errors.New("conversion table exhausted")
		}

		// Extrapolate when the date in question when we're in-between table entries, meaning:
		// - The current entry is earlier than our target (by definition true, the table is sorted)
		// - The next-up entry is later than our target
		// Otherwise, just loop on.
		if d.Type == Gregorian {
			ourNext = ConversionTable[i+1].GDate
		} else {
			ourNext = ConversionTable[i+1].JDate
		}

		af, err := d.Before(ourNext)
		if err != nil {
			return d, err
		}
		if af {
			return d.extrapolate(our, their)
		}
	}

	return d, errors.New("failed to solve conversion")
}

func (d Date) extrapolate(our, their Date) (Date, error) {
	for {
		eq, err := d.Equal(our)
		if err != nil {
			return d, err
		}
		if eq {
			return their, nil
		}
		our.Advance()
		their.Advance()
	}
}

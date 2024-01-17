package jdcal

import (
	"errors"
	"fmt"
)

const (
	maxExtrapolations = 100000
)

var (
// Variables relevant to Julian Day conversions, see
// https://aa.usno.navy.mil/downloads/c15_usb_online.pdf section 15.11.1
// Only the relevant constants for Julian and Gregorian calendars are used.
// When a constant is a conversionConstant, then the Gregorian value comes first.
// Uncomment when trying to convert via convertFromJD().

// jConstant          = 1401
// mConstant          = 2
// nConstant          = 12
// pConstant          = 1461
// qConstant          = 0
// rConstant          = 4
// sConstant          = 153
// tConstant          = 2
// uConstant          = 5
// vConstant          = 3
// wConstant          = 2
// yConstant          = 4716
// gregorianAConstant = 184
// gregorianBConstant = 274277
// gregorianCConstant = -38
)

/*
Convert converts a jdcal.Date to the "other" format: from Julian to Gregorian, or vv.
Example:

	jd, err := jdcal.New(1712, time.February, 19, jdcal.Julian)
	if err != nil {...}
	gd, err := jd.Convert()
	if err != nil {...}
	fmt.Println(gd) // Gregorian 1712/03/01
*/
func (d Date) Convert() (Date, error) {
	return d.convertFromTable()

	// convertFromJD is not applicable to the whole range of years that I want to cover.
	// Such a bummer..
	// return d.convertFromJD()
}

// func (d Date) convertFromJD() (Date, error) {
// 	// Date to JD
// 	h := int(d.Month) - mConstant
// 	g := int(d.Year) + yConstant - (nConstant-h)/nConstant
// 	f := (h - 1 + nConstant) % nConstant
// 	e := (pConstant*g+qConstant)/rConstant + d.Day - 1 - jConstant
// 	J := e + (sConstant*f+tConstant)/uConstant
// 	if d.Type == Gregorian {
// 		J = J - (3*((g+gregorianAConstant)/100))/4 - gregorianCConstant
// 	}

// 	// JD to the other Date format
// 	ret := Date{}
// 	if d.Type == Julian {
// 		ret.Type = Gregorian
// 	} else {
// 		ret.Type = Julian
// 	}

// 	f = int(J) + jConstant
// 	if ret.Type == Gregorian {
// 		f = f + (((4*J+gregorianBConstant)/146097)*3)/4 + gregorianCConstant
// 	}
// 	e = rConstant*f + vConstant
// 	g = (e % pConstant) / rConstant
// 	h = uConstant*g + wConstant

// 	ret.Day = (h%sConstant)/uConstant + 1
// 	ret.Month = time.Month((h/sConstant+mConstant)%nConstant + 1)
// 	ret.Year = Year(e/pConstant - yConstant + (nConstant+mConstant+int(ret.Month))/nConstant)
// 	return ret.adjustJD(), nil
// }

// func (d Date) adjustJD() Date {
// 	d.Year -= 1
// 	return d
// }

// convertFromTable is one implementation.
func (d Date) convertFromTable() (Date, error) {
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
	steps := 0
	startOur := our
	startTheir := their
	for {
		// Catch internal fubars.
		steps++
		if steps > maxExtrapolations {
			return d, fmt.Errorf("failed to extrapolate %v (%v to %v), %d extrapolations exceeded",
				d, startOur, startTheir, maxExtrapolations)
		}

		eq, err := d.Equal(our)
		if err != nil {
			return d, err
		}
		if eq {
			return their, nil
		}
		our = our.Forward()
		their = their.Forward()
	}
}

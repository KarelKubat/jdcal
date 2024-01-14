package jdcal

type doubleEntry struct {
	jDateStart Date
	jDateEnd   Date
	gDateStart Date
	gDateEnd   Date
	rangeType  Type
}

/*
InZone returns true when the date in question matches in the calendar progression as defined by the ZoneEntry. For example, assume a ZoneEntry:

	{
		Name: "Netherlands (Groningen City)",
		Cutovers: []Date{
			{Year: -500, Month: time.February, Day: 28, Type: Gregorian}  // Julian calendar starts
			{Year: 1583, Month: time.January, Day: 1, Type: Julian},      // switched to Gregorian
			{Year: 1594, Month: time.November, Day: 10, Type: Gregorian}, // switched to Julian
			{Year: 1700, Month: time.December, Day: 31, Type: Julian},    // switched to Gregorian
		},
	}

Stated differently, the zone is defined as follows:

	The zone calendar started on  Gregorian -500/02/28, the next day was Julian date
	Julian    1583/01/01  became  Gregorian 1583/01/11, the next day was a Gregorian date
	Gregorian 1594/11/10  became  Julian    1594/10/31, the next day was a Julian date
	Julian    1700/12/31  became  Gregorian 1701/01/11, the next day was a Gregorian date

Regarding which dates are possible in the zone, the following applies:

	1580/01/01 occurs in the zone as a Julian date, but not as as a Gregorian
	1590/01/01 occurs in the zone as a Gregorian date, but not as a Julian
	1600/01/01 occurs in the zone as a Julian date, but not as a Gregorian
	1800/01/01 occurs in the zone as a Gregorian date, but not as a Julian

Example:

	jd, err := jdcal.New(1580, 1, 1, jdcal.Julian)
	if err != nil { ... }
	in, err := jd.InZone(zoneEntry)  // obtained using jdcal.ZonesByname("Groningen City")
	if err != nil { ... }
	fmt.Println(in)                 // true

	gd, err := jdcal.New(1580, 1, 1, jdcal.Gregorian)
	if err != nil { ... }
	in, err := gd.InZone(zoneEntry)
	if err != nil { ... }
	fmt.Println(in)                 // false

Just around the cutover dates, the following applies. Around the cutover from Gregorian 1594/11/10 to Julian 1594/10/31:

- 1594/11/09 can be both a Gregorian and a Julian date. Gregorian, because it's one day before the switch over. But it can also be a Julian date, when it points 9 days beyond this switch over.

- 1594/11/10 can be both a Julian and a Gregorian date. Gregorian, because it's the switch over date. Julian, because it points 10 days beyond the switch over.

- 1594/11/11 can only be a Julian date; it doesn't exist in the Gregorian calendar in this zone, but points to the Julian calendar 11 days beyond the switch over.

Around the cutover from Julian 1700/12/31 to Gregorian Gregorian 1701/01/11, the following applies:

- 1700/12/30 and 1700/12/31 must be a Julian dates, there is no Gregorian representation. Gregorian 1700/12/30 would mean Julian 1700/12/19, and that's before the cutover.

- 1701/01/01 can't be either, it's a lost date. In the zone, the Julian calendar ends on 1700/12/31 but the Gregorian only starts on 1594/10/31.
*/
func (d Date) InZone(z ZoneEntry) (bool, error) {
	// When the date is an exact cutover, then it occurs on both types.
	for _, c := range z.Cutovers {
		// Exact dates are a match.
		eq, err := equalWithoutType(d, c)
		if err != nil {
			return false, err
		}
		if eq {
			return true, nil
		}
	}

	// Enrich the zone entry for both known date types.
	zones := make([]doubleEntry, len(z.Cutovers))
	var err error
	for i, c := range z.Cutovers {
		de := doubleEntry{}
		de.jDateStart, de.gDateStart, de.rangeType, err = datePair(c)
		if err != nil {
			return false, err
		}
		if i < len(z.Cutovers)-1 {
			de.jDateEnd, de.gDateEnd, _, err = datePair(z.Cutovers[i+1])
			if err != nil {
				return false, err
			}
		}
		zones[i] = de
	}

	// for _, z := range zones {
	// 	fmt.Println(z.jDateStart, "to", z.jDateEnd, " // ",
	// 		z.gDateStart, "to", z.gDateEnd, "type:", z.rangeType)
	// }

	// Examine the double entries, take only the types that we are interested in.
	// If the date occurs between the from/to dates, then it's a hit.
	for _, de := range zones {
		if de.rangeType != d.Type {
			continue
		}
		var start, end Date
		if de.rangeType == Julian {
			start = de.jDateStart
			end = de.jDateEnd
		} else {
			start = de.gDateStart
			end = de.gDateEnd
		}
		af, err := d.After(start)
		if err != nil {
			return false, err
		}
		var bf bool
		if !end.IsSet() {
			bf = true
		} else {
			var err error

			bf, err = d.Before(end)
			if err != nil {
				return false, err
			}
		}
		if af && bf {
			return true, nil
		}
	}
	return false, nil
}

func beforeWithoutType(a, b Date) bool {
	if a.Year < b.Year {
		return true
	}
	if a.Year > b.Year {
		return false
	}
	if a.Month < b.Month {
		return true
	}
	if a.Month > b.Month {
		return false
	}
	if a.Day < b.Day {
		return true
	}
	if a.Day > b.Day {
		return false
	}
	return false
}

func equalWithoutType(a, b Date) (bool, error) {
	if a.Type == b.Type {
		return a.Equal(b)
	}
	aa, err := a.Convert()
	if err != nil {
		return false, err
	}
	return aa.Equal(b)
}

func datePair(d Date) (jd Date, gd Date, tp Type, err error) {
	if d.Type == Julian {
		jd = d
		tp = Gregorian
		gd, err = d.Convert()
	} else {
		gd = d
		tp = Julian
		jd, err = d.Convert()
	}
	return jd, gd, tp, err
}

package jdcal

import "fmt"

/*
String returns a human-readable representation of a ZoneEntry. The returned string is not usable for machine parsing, it is only meant for human consumption. Example:

	Switzerland (Appenzell Ausserrhoden)
	  Started using the Julian    calendar   on   Gregorian -0500/02/28
	  Switched to   the Gregorian calendar   on   Julian 1584/01/01
	  Switched to   the Julian    calendar   on   Gregorian 1597/01/01
	  Switched to   the Gregorian calendar   on   Julian 1798/12/25

The first line is the start of the calendar. This will in most cases be the "start of time recording", except e.g., China where before any Gregorian calendar was used, dates would be in a calendar format that this code doesn't understand. China will have "Started using the Gregorian calendar on ..." and nothing else:

	China
	  Started using the Gregorian calendar   on   Julian 1911/12/01

Next entries, when present, are switches. In the example of Switzerland (Appenzell Ausserrhoden), the zone switched to the Gregorian calendar on 1584/01/01 (which must be a Julian date, because they were switching). Then they switched back to Julian on 1597/01/01 (on the Gregorian calendar, because they were switching). Etc..
*/
func (z ZoneEntry) String() string {
	out := z.Name
	for i, c := range z.Cutovers {
		var change string
		if i == 0 {
			change = "Started using"
		} else {
			change = "Switched to  "
		}
		var otherType Type
		if c.Type == Gregorian {
			otherType = Julian
		} else {
			otherType = Gregorian
		}
		out += fmt.Sprintf(`
  %s the %-9s calendar   on   %v`, change, otherType, c)
	}
	return out
}

package jdcal

import "strings"

/*
ZonesByname returns a list of ZoneEntry's matching the input argument. The matching is done without regard to case. E.g., ZonesByName("netherlands") will return zones for "Belgium (Southern Netherlands)", "Netherlands (Brabant)" and ~7 more.

Example:

	fmt.Println(jdcal.ZonesByName("netherlands"))

	// Output:
	// Belgium (Southern Netherlands)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1582/12/20
	// Netherlands (Brabant)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1582/12/14
	// Netherlands (Drenthe)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1701/04/30
	// Netherlands (Frisia)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1701/12/31
	// Netherlands (Gelderland)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1700/06/12
	// Netherlands (Groningen City)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1583/01/01
	//   Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	//   Switched to   the Gregorian calendar   on   Julian 1700/12/31
	// Netherlands (Holland)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1583/01/01
	// Netherlands (Utrecht, Overijssel)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1700/11/30
	// Netherlands (Zeeland, States General)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1582/12/14
*/
func ZonesByName(n string) []ZoneEntry {
	n = strings.ToLower(n)
	ret := []ZoneEntry{}
	for _, z := range ZonesTable {
		if strings.Contains(strings.ToLower(z.Name), n) {
			ret = append(ret, z)
		}
	}
	return ret
}

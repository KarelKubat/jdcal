package jdcal

import "strings"

/*
ZonesByname returns a list of ZoneEntry's matching the input argument. The matching is done without regard to case. E.g., ZonesByName("netherlands") will return zones for "Belgium (Southern Netherlands)", "Netherlands (Brabant)" and ~7 more.
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

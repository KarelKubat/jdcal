package jdcal

import "strings"

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

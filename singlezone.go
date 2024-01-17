package jdcal

import (
	"fmt"
	"strings"
)

/*
SingleZone returns a ZoneEntry matching a name, or an error when nothing matches, or when multiple zones match.

Example:

	var zn jdcal.ZoneEntry
	var err error

	zn, err = jdcal.SingleZone("netherlands")  // error: >1 matches
	zn, err = jdcal.SingleZone("gelderland")   // nil error, zn is valid
	zn, err = jdcal.SingleZone("xyzzy")        // error: 0 matches
*/
func SingleZone(n string) (z ZoneEntry, err error) {
	zones := ZonesByName(n)

	if len(zones) == 0 {
		return ZoneEntry{}, fmt.Errorf("there is no zone matching %q", n)
	}
	if len(zones) > 1 {
		names := []string{}
		for _, z := range zones {
			names = append(names, z.Name)
		}
		return ZoneEntry{}, fmt.Errorf("multiple zones match %q: %s", n, strings.Join(names, " / "))
	}

	return zones[0], nil
}

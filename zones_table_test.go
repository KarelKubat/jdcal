package jdcal

import (
	"testing"
)

func TestZonesTable(t *testing.T) {

	// TODO:
	// Ensure that cutover entries are chronoligically increasing,
	// ZonesForDate e.a. depend on it.
	//
	if len(ZonesTable) == 0 {
		t.Fatalf("ZonesTable is empty")
	}

	namesSeen := map[string]struct{}{}
	cutoversSeen := map[string]string{}
	for i, e := range ZonesTable {
		// Avoid double names
		if _, ok := namesSeen[e.Name]; ok {
			t.Errorf("FAIL: ZonesTable[%d]: duplicate name %q", i, e.Name)
		}
		namesSeen[e.Name] = struct{}{}

		// First entry most often is of the type Gregorian, i.e, on this date the Julian calendar
		// went into effect.
		if e.Cutovers[0].Type != Gregorian {
			t.Logf("%q: first entry %v doesn't indicate starting on the Julian calendar\n",
				e.Name, e.Cutovers[0])
		}

		// Last entry must be of type Julian to indicate a swap into the Gregorian calendar.
		if e.Cutovers[len(e.Cutovers)-1].Type != Julian {
			t.Errorf("FAIL: %q: last entry %v doesn't indicate starting on the Gregorian calendar\n",
				e.Name, e.Cutovers[len(e.Cutovers)-1])
		}

		// Check for folding
		cutovers := ""
		for _, c := range e.Cutovers {
			cutovers += c.String()
		}
		if otherZone, ok := cutoversSeen[cutovers]; ok {
			t.Logf("%q could be folded with %q", e.Name, otherZone)
		}
		cutoversSeen[cutovers] = e.Name
	}
}

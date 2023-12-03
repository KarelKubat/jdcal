package jdcal

import "testing"

func TestLast(t *testing.T) {
	// The table ends around 2000. This is just a rough test to catch big table errors.
	for _, tp := range []Type{Gregorian, Julian} {
		d := Last(tp)
		if d.Year < 1900 {
			t.Errorf("Last(%v).Year = %v, need >=600", tp, d.Year)
		}
	}
}

package jdcal

import (
	"testing"
)

func TestFirst(t *testing.T) {
	// The table starts around -500. This is just a rough test to catch big table errors.
	for _, tp := range []Type{Gregorian, Julian} {
		d := First(tp)
		if d.Year < -600 {
			t.Errorf("First(%v).Year = %v, need <=600", tp, d.Year)
		}
	}
}

package jdcal

import "testing"

func TestTable(t *testing.T) {
	// Test that the table has valid dates.
	for _, e := range Table {
		for _, d := range []Date{e.JDate, e.GDate} {
			if err := d.Valid(); err != nil {
				t.Errorf("%+v .Valid() = %q, want nil error", d, err.Error())
			}
		}
	}
}

package jdcal

import "testing"

func TestConversionTable(t *testing.T) {
	// Test that the table has valid dates.

	for _, e := range ConversionTable {
		check := func(err error) {
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}
		for _, d := range []Date{e.JDate, e.GDate} {
			bf, err := d.Before(First(d.Type))
			check(err)
			af, err := d.After(First(d.Type))
			check(err)

			if !bf && !af {
				if err := d.Valid(); err != nil {
					t.Errorf("%+v .Valid() = %q, want nil error", d, err.Error())
				}
			}
		}
	}
}

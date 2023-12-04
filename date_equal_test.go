package jdcal

import (
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		dt        Date
		ot        Date
		wantError bool
		wantEqual bool
	}{
		{
			dt:        Date{Year: 1962, Month: time.August, Day: 19, Type: Gregorian},
			ot:        Date{Year: 1962, Month: time.August, Day: 19, Type: Gregorian},
			wantError: false,
			wantEqual: true,
		},
		{
			dt:        Date{Year: 1962, Month: time.August, Day: 19, Type: Gregorian},
			ot:        Date{Year: 1962, Month: time.August, Day: 18, Type: Gregorian},
			wantError: false,
			wantEqual: false,
		},
		{
			dt:        Date{Year: 1962, Month: time.August, Day: 19, Type: Gregorian},
			ot:        Date{Year: 1962, Month: time.August, Day: 20, Type: Gregorian},
			wantError: false,
			wantEqual: false,
		},
		{
			dt:        Date{Year: 1962, Month: time.August, Day: 19, Type: Gregorian},
			ot:        Date{Year: 1962, Month: time.August, Day: 20, Type: Julian},
			wantError: true,
			wantEqual: false,
		},
	} {
		for _, d := range []Date{test.dt, test.ot} {
			if err := d.Valid(); err != nil {
				t.Fatalf("%+v .Valid() = %q, need nil error", test.dt, err.Error())
			}
		}
		gotEqual, err := test.dt.Equal(test.ot)
		gotError := err != nil
		switch {
		case gotError != test.wantError:
			t.Fatalf("%+v Equal(%+v) = _,%v, want error: %v", test.dt, test.ot, err, test.wantError)
		case gotEqual != test.wantEqual:
			t.Errorf("%+v Equal(%+v) = %v,_, want %v", test.dt, test.ot, gotEqual, test.wantEqual)
		}
	}
}

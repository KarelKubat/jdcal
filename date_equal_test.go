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
		d, err := New(test.dt)
		if err != nil {
			t.Fatalf("New(%v) = _,%q, need nil error", test.dt, err.Error())
		}
		o, err := New(test.ot)
		if err != nil {
			t.Fatalf("New(%v) = _,%q, need nil error", test.dt, err.Error())
		}
		gotEqual, err := d.Equal(o)
		gotError := err != nil
		switch {
		case gotError != test.wantError:
			t.Fatalf("%+v Equal(%+v) = _,%v, want error: %v", d, o, err, test.wantError)
		case gotEqual != test.wantEqual:
			t.Errorf("%+v Equal(%+v) = %v,_, want %v", d, o, gotEqual, test.wantEqual)
		}
	}
}

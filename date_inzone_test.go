package jdcal

import (
	"testing"
	"time"
)

func TestDateInzone(t *testing.T) {
	z := ZoneEntry{
		Name: "testzone",
		Cutovers: []Date{
			{Year: -500, Month: time.February, Day: 28, Type: Gregorian},
			{Year: 1584, Month: time.January, Day: 1, Type: Julian},
			{Year: 1597, Month: time.January, Day: 1, Type: Gregorian},
			{Year: 1798, Month: time.December, Day: 25, Type: Julian},
		},
	}

	for _, test := range []struct {
		d       Date
		wantIn  bool
		wantErr bool
	}{
		{
			d:       Date{Year: 1000, Month: time.January, Day: 1, Type: Julian},
			wantIn:  true,
			wantErr: false,
		},
	} {
		in, err := test.d.InZone(z)
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("%+v .Inzone(%v) = _,%q, want nil error", test.d, z, err.Error())
		case err == nil && test.wantErr:
			t.Errorf("%+v .Inzone(%v) = _,nil, want error", test.d, z)
		case err == nil && !test.wantErr && in != test.wantIn:
			t.Errorf("%+v .Inzone(%v) = %v,_, want %v", test.d, z, in, test.wantIn)
		}
	}
}

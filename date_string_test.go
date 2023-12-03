package jdcal

import (
	"testing"
	"time"
)

func TestDateString(t *testing.T) {
	for _, test := range []struct {
		d          Date
		wantString string
	}{
		{
			d:          Date{Year: 1600, Month: time.August, Day: 19, Type: Gregorian},
			wantString: "Gregorian 1600/08/19",
		},
		{
			d:          Date{Year: 1600, Month: time.August, Day: 19, Type: Julian},
			wantString: "Julian 1600/08/19",
		},
		{
			d:          Date{Year: -500, Month: time.August, Day: 19, Type: Gregorian},
			wantString: "Gregorian -0500/08/19",
		},
		{
			d:          Date{Year: -500, Month: time.August, Day: 19, Type: Julian},
			wantString: "Julian -0500/08/19",
		},
	} {
		if gotString := test.d.String(); gotString != test.wantString {
			t.Errorf("%+v .String() = %q, want %q", test.d, gotString, test.wantString)
		}
	}
}

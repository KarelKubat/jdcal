package jdcal

import (
	"testing"
	"time"
)

func TestDateAfterOrEqual(t *testing.T) {
	d := Date{Year: 1962, Month: time.August, Day: 19}
	for _, test := range []struct {
		other     Date
		wantErr   bool
		wantAfter bool
	}{
		{
			other:     Date{Year: 1962, Month: time.August, Day: 18},
			wantAfter: true,
		},
		{
			other:     Date{Year: 1962, Month: time.August, Day: 19},
			wantAfter: true,
		},
		{
			other:     Date{Year: 1962, Month: time.August, Day: 20},
			wantAfter: false,
		},
		{
			other:   Date{Year: 1962, Month: time.August, Day: 20, Type: Julian},
			wantErr: true,
		},
	} {
		gotAfter, err := d.AfterOrEqual(test.other)
		gotErr := err != nil
		switch {
		case gotErr != test.wantErr:
			t.Errorf("%+v .AfterOrEqual(%v): gotErr=%v, wantErr=%v", d, test.other, gotErr, test.wantErr)
		case !gotErr && gotAfter != test.wantAfter:
			t.Errorf("%+v .AfterOrEqual(%v) = %v,_, want %v", d, test.other, gotAfter, test.wantAfter)
		}
	}
}

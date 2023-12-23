package jdcal

import (
	"testing"
	"time"
)

func TestDateBeforeEqual(t *testing.T) {
	d := Date{Year: 1962, Month: time.August, Day: 19}
	for _, test := range []struct {
		other      Date
		wantErr    bool
		wantBefore bool
	}{
		{
			other:      Date{Year: 1962, Month: time.August, Day: 18},
			wantBefore: false,
		},
		{
			other:      Date{Year: 1962, Month: time.August, Day: 19},
			wantBefore: true,
		},
		{
			other:      Date{Year: 1962, Month: time.August, Day: 20},
			wantBefore: true,
		},
		{
			other:   Date{Year: 1962, Month: time.August, Day: 20, Type: Julian},
			wantErr: true,
		},
	} {
		gotBefore, err := d.BeforeOrEqual(test.other)
		gotErr := err != nil
		switch {
		case gotErr != test.wantErr:
			t.Errorf("%+v .BeforeOrEqual(%v): gotErr=%v, wantErr=%v", d, test.other, gotErr, test.wantErr)
		case !gotErr && gotBefore != test.wantBefore:
			t.Errorf("%+v .BeforeOrEqual(%v) = %v,_, want %v", d, test.other, gotBefore, test.wantBefore)
		}
	}
}

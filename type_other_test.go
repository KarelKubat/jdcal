package jdcal

import "testing"

func TestOtherType(t *testing.T) {
	for _, test := range []struct {
		tp        Type
		wantOther Type
	}{
		{
			tp:        Gregorian,
			wantOther: Julian,
		},
		{
			tp:        Julian,
			wantOther: Gregorian,
		},
	} {
		if got := test.tp.Other(); got != test.wantOther {
			t.Errorf("%v .Other() = %v, want %v", test.tp, got, test.wantOther)
		}
	}
}

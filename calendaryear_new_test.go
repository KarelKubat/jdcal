package jdcal

import (
	"testing"
)

func TestNewCalendarYear(t *testing.T) {
	for _, test := range []struct {
		yr      Year
		wantErr bool
	}{
		{
			yr:      -999999,
			wantErr: true,
		},
		{
			yr:      0,
			wantErr: true,
		},
		{
			yr:      999999,
			wantErr: true,
		},
		{
			yr:      1,
			wantErr: false,
		},
	} {
		_, err := NewCalendarYear(test.yr, Julian)
		if gotErr := err != nil; gotErr != test.wantErr {
			t.Errorf("NewCalendarYear(%v,_) = _,%v; got error: %v, want error: %v", test.yr, err, gotErr, test.wantErr)
		}
	}
}

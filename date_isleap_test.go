package jdcal

import (
	"testing"
	"time"
)

func TestIsLeap(t *testing.T) {
	for _, test := range []struct {
		tp         Type
		year       Year
		wantIsLeap bool
	}{
		// For Julian dates it's just "Divisible by 4? Then yes."
		{
			tp:         Julian,
			year:       1003,
			wantIsLeap: false,
		},
		{
			tp:         Julian,
			year:       1004,
			wantIsLeap: true,
		},
		{
			tp:         Julian,
			year:       1600,
			wantIsLeap: true,
		},
		{
			tp:         Julian,
			year:       1900,
			wantIsLeap: true,
		},
		{
			tp:         Julian,
			year:       2000,
			wantIsLeap: true,
		},
		// For Gregorian dates it's "Divisible by 4? Then yes, except when it's a century.
		// Then it must be divisible by 400."
		{
			tp:         Gregorian,
			year:       1003,
			wantIsLeap: false,
		},
		{
			tp:         Gregorian,
			year:       1004,
			wantIsLeap: true,
		},
		{
			tp:         Gregorian,
			year:       1600,
			wantIsLeap: true,
		},
		{
			tp:         Gregorian,
			year:       1900,
			wantIsLeap: false,
		},
		{
			tp:         Gregorian,
			year:       2000,
			wantIsLeap: true,
		},
		// Negative (BC); -1, -5, -9 etc. are leap years
		{
			tp:         Julian,
			year:       -400,
			wantIsLeap: false,
		},
		{
			tp:         Julian,
			year:       -401,
			wantIsLeap: true,
		},
		{
			tp:         Julian,
			year:       -400,
			wantIsLeap: false,
		},
	} {
		d := Date{Year: test.year, Month: time.January, Day: 1, Type: test.tp}
		if gotIsLeap := d.IsLeap(); gotIsLeap != test.wantIsLeap {
			t.Errorf("%+v .IsLeap() = %v, want %v", d, gotIsLeap, test.wantIsLeap)
		}
	}
}

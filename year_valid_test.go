package jdcal

import "testing"

func TestYearValid(t *testing.T) {
	for _, test := range []struct {
		year    Year
		wantErr bool
	}{
		{
			// Before conversion table
			year:    -2000,
			wantErr: true,
		},
		{
			// After conversion table
			year:    20000,
			wantErr: true,
		},
		{
			// Zero
			year:    0,
			wantErr: true,
		},
		{
			// Valid
			year:    1,
			wantErr: false,
		},
	} {
		err := test.year.Valid()
		gotErr := err != nil
		if gotErr != test.wantErr {
			t.Errorf("%v.Valid(): goterr=%v, wanterr=%v", test.year, gotErr, test.wantErr)
		}
	}
}

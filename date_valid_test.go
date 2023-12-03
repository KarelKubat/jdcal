package jdcal

import (
	"strings"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	for _, test := range []struct {
		d       Date
		wantErr string
	}{
		{
			d:       Date{Year: 1962, Month: time.August, Day: 19},
			wantErr: "",
		},
		{
			d:       Date{Year: 1962, Month: time.August, Day: 32},
			wantErr: "can't have 32 days",
		},
		{
			d:       Date{Year: -1962, Month: time.August, Day: 19},
			wantErr: "is before the first convertible date",
		},
		{
			d:       Date{Year: 11962, Month: time.August, Day: 19},
			wantErr: "is after the last convertible date",
		},
		{
			d:       Date{Year: 1900, Month: time.February, Day: 29},
			wantErr: "1900 has no February 29th",
		},
		{
			d:       Date{Year: 2000, Month: time.February, Day: 29},
			wantErr: "",
		},
	} {
		err := test.d.Valid()
		switch {
		case err == nil && test.wantErr != "":
			t.Errorf("%+v .Valid() = nil, want error with %q", test.d, test.wantErr)
		case err != nil && test.wantErr == "":
			t.Errorf("%+v .Valid() = %q, want nil error", test.d, err)
		case err != nil && test.wantErr != "" && !strings.Contains(err.Error(), test.wantErr):
			t.Errorf("%+v .Valid() = %q, want error with %q", test.d, err, test.wantErr)
		}
	}
}

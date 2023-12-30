package jdcal

import (
	"strings"
	"testing"
	"time"
)

func TestStringToYMD(t *testing.T) {
	for _, test := range []struct {
		arg       string
		wantErr   string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{
			arg:     "one",
			wantErr: "malformed",
		},
		{
			arg:     "one/two",
			wantErr: "malformed",
		},
		{
			arg:     "one/two/three/four",
			wantErr: "malformed",
		},
		{
			arg:       "1234/05/06",
			wantErr:   "",
			wantYear:  1234,
			wantMonth: time.May,
			wantDay:   6,
		},
		{
			arg:     "bla/05/06",
			wantErr: "year part",
		},
		{
			arg:     "1234/bla/06",
			wantErr: "month part",
		},
		{
			arg:     "1234/05/bla",
			wantErr: "day part",
		},
	} {
		y, m, d, err := StringToYMD(test.arg)
		if err != nil {
			if test.wantErr == "" {
				t.Errorf("StringToYMD(%q) = _,_,_,%q, want nil error", test.arg, err.Error())
			} else if !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("StringToYMD(%q) = _,_,_,%q, want error with %q",
					test.arg, err.Error(), test.wantErr)
			}
			continue
		}
		if y != test.wantYear || m != test.wantMonth || d != test.wantDay {
			t.Errorf("StringToYMD(%q) = %v,%v,%v, want %v,%v,%v",
				test.arg, y, m, d, test.wantYear, test.wantMonth, test.wantDay)
		}
	}
}

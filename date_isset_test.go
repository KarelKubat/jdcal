package jdcal

import "testing"

func TestIsSet(t *testing.T) {
	for _, test := range []struct {
		d    Date
		want bool
	}{
		{
			d:    Date{Year: 0, Month: 0, Day: 0},
			want: false,
		},
		{
			d:    Date{Year: 1, Month: 0, Day: 0},
			want: true,
		},
		{
			d:    Date{Year: 0, Month: 1, Day: 0},
			want: true,
		},
		{
			d:    Date{Year: 0, Month: 0, Day: 1},
			want: true,
		},
	} {
		if got := test.d.IsSet(); got != test.want {
			t.Errorf("%+v .IsSet() = %v, want %v", test.d, got, test.want)
		}
	}
}

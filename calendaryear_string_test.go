package jdcal

import "testing"

func TestCalendarYearString(t *testing.T) {
	for _, test := range []struct {
		yr   Year
		tp   Type
		want string
	}{
		// 4 full digits
		{
			yr:   1600,
			tp:   Julian,
			want: "Julian 1600",
		},
		{
			yr:   1600,
			tp:   Gregorian,
			want: "Gregorian 1600",
		},
		// Zero padding
		{
			yr:   16,
			tp:   Julian,
			want: "Julian 0016",
		},
		{
			yr:   16,
			tp:   Gregorian,
			want: "Gregorian 0016",
		},
		// 4 full digits negative
		{
			yr:   -1600,
			tp:   Julian,
			want: "Julian -1600",
		},
		{
			yr:   -1600,
			tp:   Gregorian,
			want: "Gregorian -1600",
		},
		// Zero padding
		{
			yr:   -16,
			tp:   Julian,
			want: "Julian -0016",
		},
		{
			yr:   -16,
			tp:   Gregorian,
			want: "Gregorian -0016",
		},
	} {
		cyr := CalendarYear{Year: test.yr, Type: test.tp}
		if got := cyr.String(); got != test.want {
			t.Errorf("%+v .String() = %q, want %q", cyr, got, test.want)
		}
	}
}

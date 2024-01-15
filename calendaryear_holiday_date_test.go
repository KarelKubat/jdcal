package jdcal

import (
	"testing"
	"time"
)

func TestHolidayDate(t *testing.T) {
	// A few canned dates.
	for _, test := range []struct {
		cyr     CalendarYear
		holiday Holiday
		want    Date
	}{
		{
			cyr:     CalendarYear{Year: 1369, Type: Julian},
			holiday: Easter,
			want:    Date{1369, time.March, 25, Julian},
		},
		{
			cyr:     CalendarYear{Year: 1370, Type: Julian},
			holiday: Easter,
			want:    Date{1370, time.April, 14, Julian},
		},
		{
			cyr:     CalendarYear{Year: 1371, Type: Julian},
			holiday: Easter,
			want:    Date{1371, time.April, 6, Julian},
		},

		{
			cyr:     CalendarYear{Year: 2023, Type: Gregorian},
			holiday: Easter,
			want:    Date{2023, time.April, 9, Gregorian},
		},
		{
			cyr:     CalendarYear{Year: 2023, Type: Gregorian},
			holiday: Ascension,
			want:    Date{2023, time.May, 18, Gregorian},
		},
		{
			cyr:     CalendarYear{Year: 2023, Type: Gregorian},
			holiday: Pentecost,
			want:    Date{2023, time.May, 28, Gregorian},
		},
	} {
		got, err := test.cyr.HolidayDate(test.holiday)
		if err != nil {
			t.Errorf("%+v .HolidayDate(%v) = _,%q, want nil error", test.cyr, test.holiday, err.Error())
			continue
		}
		eq, err := got.Equal(test.want)
		if err != nil {
			t.Errorf("%+v .Equal(%+v) = _,%q, want nil error", got, test.want, err.Error())
			continue
		}
		if !eq {
			t.Errorf("%+v .HolidayDate(%v) = %v,_, want %v", test.cyr, test.holiday, got, test.want)
		}
	}
}

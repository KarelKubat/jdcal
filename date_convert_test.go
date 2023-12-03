package jdcal

import (
	"testing"
	"time"
)

func TestConvertSome(t *testing.T) {
	for _, test := range []struct {
		desc string
		jd   Date
		gd   Date
	}{
		{
			desc: "start of epoch (lookup hit)",
			jd:   First(Julian),
			gd:   First(Gregorian),
		},
		{
			desc: "end of epoch (lookup hit)",
			jd:   Last(Julian),
			gd:   Last(Gregorian),
		},
		{
			desc: "table middle lookup hit",
			jd:   Date{Year: 1500, Month: time.February, Day: 29, Type: Julian},
			gd:   Date{Year: 1500, Month: time.March, Day: 10, Type: Gregorian},
		},
		{
			desc: "outside of table hit (1)",
			jd:   Date{Year: 1700, Month: time.March, Day: 2, Type: Julian},
			gd:   Date{Year: 1700, Month: time.March, Day: 13, Type: Gregorian},
		},
		{
			desc: "outside of table hit (2)",
			jd:   Date{Year: 1700, Month: time.March, Day: 3, Type: Julian},
			gd:   Date{Year: 1700, Month: time.March, Day: 14, Type: Gregorian},
		},
		{
			desc: "outside of table hit (3)",
			jd:   Date{Year: 1700, Month: time.March, Day: 4, Type: Julian},
			gd:   Date{Year: 1700, Month: time.March, Day: 15, Type: Gregorian},
		},
		{
			desc: "outside of table hit (4)",
			jd:   Date{Year: 1700, Month: time.March, Day: 5, Type: Julian},
			gd:   Date{Year: 1700, Month: time.March, Day: 16, Type: Gregorian},
		},
	} {
		run := func(desc string, a, b Date) {
			gotOut, err := a.Convert()
			if err != nil {
				t.Fatalf("%q: %+v .Convert() = _,%q, need nil error", desc, a, err.Error())
			}
			eq, err := b.Equal(gotOut)
			if err != nil {
				t.Fatalf("%q: %+v .Equal(%v) = _,%q, need nil error", desc, a, gotOut, err.Error())
			}
			if !eq {
				t.Errorf("%q: %+v .Convert() = %v, want %v", desc, a, gotOut, b)
			}
		}
		run(test.desc, test.jd, test.gd)
		run(test.desc, test.gd, test.jd)
	}
}

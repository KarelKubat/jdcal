package jdcal

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"

	"github.com/KarelKubat/jdcal/main/bigconversiontest/testdates"
)

const (
	// # of conversions to pick from main/bigconversiontest/testdates/testdates.go for testing
	// Should be a large number
	nConversionTests = 100
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

var diffJulianToGregorian = map[Year][]int{}
var diffGregorianToJulian = map[Year][]int{}

func logYearDiff(mp *map[Year][]int, y Year, diff int) {
	_, ok := (*mp)[y]
	if !ok {
		(*mp)[y] = []int{diff}
	}
	if !slices.Contains((*mp)[y], diff) {
		(*mp)[y] = append((*mp)[y], diff)
	}
}

func reportDiffs(t *testing.T, mp map[Year][]int, title string) {
	if len(mp) == 0 {
		return
	}
	t.Errorf("*** DIFFS IN YEARS for %s ***", title)
	yrs := []Year{}
	for y := range mp {
		yrs = append(yrs, y)
	}
	slices.Sort(yrs)
	for _, y := range yrs {
		t.Errorf("%v: %v", y, mp[y])
	}
}

func TestConvertFromFullSet(t *testing.T) {
	for i := 0; i < nConversionTests; i++ {
		if err := oneTest(); err != nil {
			t.Error(err)
		}
	}
	reportDiffs(t, diffJulianToGregorian, "Julian to Gregorian")
	reportDiffs(t, diffGregorianToJulian, "Gregorian to Julian")
}

func oneTest() error {
	test := testdates.TestDates[rand.Intn(len(testdates.TestDates))]
	jd := Date{Year: Year(test.J.Year), Month: test.J.Month, Day: test.J.Day, Type: Julian}
	gd := Date{Year: Year(test.G.Year), Month: test.G.Month, Day: test.G.Day, Type: Gregorian}

	gd1, err := jd.Convert()
	if err != nil {
		return fmt.Errorf("%+v .Convert() = _,%q, need nil error", jd, err.Error())
	}
	eq, err := gd.Equal(gd1)
	if err != nil {
		return fmt.Errorf("%+v .Equal(%+v) = _,%q, need nil error", gd, gd1, err.Error())
	}
	if !eq {
		logYearDiff(&diffJulianToGregorian, jd.Year, int(jd.Year-gd1.Year))
		return fmt.Errorf("reference %+v .Equal(computed %+v) = false, want true", gd, gd1)
	}

	jd1, err := gd.Convert()
	if err != nil {
		return fmt.Errorf("%+v .Convert() = _,%q, need nil error", jd1, err.Error())
	}
	eq, err = jd.Equal(jd1)
	if err != nil {
		return fmt.Errorf("%+v .Equal(%+v) = _,%q, need nil error", jd, jd1, err.Error())
	}
	if !eq {
		logYearDiff(&diffGregorianToJulian, gd.Year, int(gd.Year-jd1.Year))
		return fmt.Errorf("reference %+v .Equal(computed %+v) = false, want true", jd, jd1)
	}
	return nil
}

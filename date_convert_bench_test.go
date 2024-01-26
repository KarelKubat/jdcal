package jdcal

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/KarelKubat/jdcal/main/bigconversiontest/testdates"
)

const (
	nBenchConversions = 10000
)

func BenchmarkDateConvert(b *testing.B) {
	check := func(err error) {
		if err != nil {
			b.Errorf("unexpected error %q", err.Error())
		}
	}

	for _, a := range []ConversionAlgorithm{ByLookup, ByProgression} {
		b.Run(fmt.Sprintf("Algorithm %s, %d random conversions", a, nBenchConversions),
			func(b *testing.B) {
				Algorithm = a
				for i := 0; i < nBenchConversions; i++ {
					test := testdates.TestDates[rand.Intn(len(testdates.TestDates))]
					jd := Date{Year: Year(test.J.Year), Month: test.J.Month, Day: test.J.Day, Type: Julian}
					gd := Date{Year: Year(test.G.Year), Month: test.G.Month, Day: test.G.Day, Type: Gregorian}

					_, err := jd.Convert()
					check(err)
					_, err = gd.Convert()
					check(err)
				}
			})
	}
}

package jdcal

import (
	"testing"
	"time"
)

func TestForward(t *testing.T) {
	for _, test := range []struct {
		desc string
		d    Date
		next []Date
	}{
		{
			desc: "most simple: forward over January, Gregorian",
			d:    Date{Year: 1962, Month: time.January, Day: 28, Type: Gregorian},
			next: []Date{
				{Year: 1962, Month: time.January, Day: 29, Type: Gregorian},
				{Year: 1962, Month: time.January, Day: 30, Type: Gregorian},
				{Year: 1962, Month: time.January, Day: 31, Type: Gregorian},
				{Year: 1962, Month: time.February, Day: 1, Type: Gregorian},
				{Year: 1962, Month: time.February, Day: 2, Type: Gregorian},
			},
		},
		{
			desc: "most simple: forward over January, Julian",
			d:    Date{Year: 1962, Month: time.January, Day: 28, Type: Julian},
			next: []Date{
				{Year: 1962, Month: time.January, Day: 29, Type: Julian},
				{Year: 1962, Month: time.January, Day: 30, Type: Julian},
				{Year: 1962, Month: time.January, Day: 31, Type: Julian},
				{Year: 1962, Month: time.February, Day: 1, Type: Julian},
				{Year: 1962, Month: time.February, Day: 2, Type: Julian},
			},
		},
		{
			desc: "forward over February, non-leap Gregorian",
			d:    Date{Year: 1962, Month: time.February, Day: 26, Type: Gregorian},
			next: []Date{
				{Year: 1962, Month: time.February, Day: 27, Type: Gregorian},
				{Year: 1962, Month: time.February, Day: 28, Type: Gregorian},
				{Year: 1962, Month: time.March, Day: 1, Type: Gregorian},
				{Year: 1962, Month: time.March, Day: 2, Type: Gregorian},
			},
		},
		{
			desc: "forward over February, century not divisible 400, Gregorian",
			d:    Date{Year: 1900, Month: time.February, Day: 26, Type: Gregorian},
			next: []Date{
				{Year: 1900, Month: time.February, Day: 27, Type: Gregorian},
				{Year: 1900, Month: time.February, Day: 28, Type: Gregorian},
				{Year: 1900, Month: time.March, Day: 1, Type: Gregorian},
				{Year: 1900, Month: time.March, Day: 2, Type: Gregorian},
			},
		},
		{
			desc: "forward over February, century divisible 400, Gregorian",
			d:    Date{Year: 2000, Month: time.February, Day: 26, Type: Gregorian},
			next: []Date{
				{Year: 2000, Month: time.February, Day: 27, Type: Gregorian},
				{Year: 2000, Month: time.February, Day: 28, Type: Gregorian},
				{Year: 2000, Month: time.February, Day: 29, Type: Gregorian},
				{Year: 2000, Month: time.March, Day: 1, Type: Gregorian},
				{Year: 2000, Month: time.March, Day: 2, Type: Gregorian},
			},
		},
		{
			desc: "forward over February, century not divisible by 400, Julian",
			d:    Date{Year: 1900, Month: time.February, Day: 26, Type: Julian},
			next: []Date{
				{Year: 1900, Month: time.February, Day: 27, Type: Julian},
				{Year: 1900, Month: time.February, Day: 28, Type: Julian},
				{Year: 1900, Month: time.February, Day: 29, Type: Julian},
				{Year: 1900, Month: time.March, Day: 1, Type: Julian},
				{Year: 1900, Month: time.March, Day: 2, Type: Julian},
			},
		},
		{
			desc: "forward over February, century divisible 400, Julian",
			d:    Date{Year: 2000, Month: time.February, Day: 26, Type: Julian},
			next: []Date{
				{Year: 2000, Month: time.February, Day: 27, Type: Julian},
				{Year: 2000, Month: time.February, Day: 28, Type: Julian},
				{Year: 2000, Month: time.February, Day: 29, Type: Julian},
				{Year: 2000, Month: time.March, Day: 1, Type: Julian},
				{Year: 2000, Month: time.March, Day: 2, Type: Julian},
			},
		},
		{
			desc: "forward over new year, Gregorian",
			d:    Date{Year: 1234, Month: time.December, Day: 30, Type: Gregorian},
			next: []Date{
				{Year: 1234, Month: time.December, Day: 31, Type: Gregorian},
				{Year: 1235, Month: time.January, Day: 1, Type: Gregorian},
				{Year: 1235, Month: time.January, Day: 2, Type: Gregorian},
			},
		},
		{
			desc: "forward over new year, Julian",
			d:    Date{Year: 1234, Month: time.December, Day: 30, Type: Julian},
			next: []Date{
				{Year: 1234, Month: time.December, Day: 31, Type: Julian},
				{Year: 1235, Month: time.January, Day: 1, Type: Julian},
				{Year: 1235, Month: time.January, Day: 2, Type: Julian},
			},
		},
		{
			desc: "forward over the non-existing year zero, Julian",
			d:    Date{Year: -1, Month: time.December, Day: 30, Type: Julian},
			next: []Date{
				{Year: -1, Month: time.December, Day: 31, Type: Julian},
				{Year: 1, Month: time.January, Day: 1, Type: Julian},
				{Year: 1, Month: time.January, Day: 2, Type: Julian},
			},
		},
	} {
		for i, nxt := range test.next {
			test.d = test.d.Forward()
			eq, err := test.d.Equal(nxt)
			if err != nil {
				t.Fatalf("%+v .Equal(%+v) = _,%q, need nil error", test.d, nxt, err.Error())
			}
			if !eq {
				t.Errorf("%q: after round %d: advanced to %v, want %v", test.desc, i+1, test.d, nxt)
			}
		}
	}
}

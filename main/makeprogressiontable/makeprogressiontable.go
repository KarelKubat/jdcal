package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/KarelKubat/jdcal"
)

/*
	Make sure that startYear matches StartProgressionYear with a correct startOffset,
	and that endYear matches EndProgressionYear!
*/

const (
	startYear   = -506 // generate from here
	endYear     = 2106 // until here
	startOffset = -6   // Gregorian is 5 days behind on startYear

	header = `package jdcal

// This is a generated file, don't edit. Look for a nearby Makefile.
	
import "time"

const (
	StartProgressionYear = -500
	EndProgressionYear   = 2100

)

/*
MonthProgression maps arrays of ordinals to months. There is such a table for leap years
(LeapMonthProgression) and for non-leap years (NonLeapMonthProgression).
*/
type MonthProgression map[time.Month][]Ordinal

`
)

var outputLines = []string{}

func main() {
	if len(os.Args) != 2 {
		check(errors.New("usage: go run main/progressiontable/progressiontable.go $OUTPUTFILE"))
	}

	outputLines = append(outputLines, header)

	var leapYearDays, nonLeapYearDays int

	leapCyr := jdcal.CalendarYear{Year: 12, Type: jdcal.Julian}
	leapDaysPerMonth := leapCyr.DaysPerMonth()
	output(`/*
LeapMonthProgression holds per month the ordinal of a day,
January 1st being 0, December 31st being 365 (365 days in a year, plus February 29th
is 366, so ordinal number 365).
*/
`)
	output("var LeapMonthProgression = MonthProgression{\n")
	count := 0
	for m := time.January; m <= time.December; m++ {
		output("time.%v: {\n", m)
		output("0, // filler\n")
		for d := 1; d <= leapDaysPerMonth[m]; d++ {
			output("%d, ", count)
			count++
			leapYearDays = count // updated every time, don't care, it's code generation
		}
		output("\n},\n")
	}
	output("}\n\n")

	nonLeapCyr := jdcal.CalendarYear{Year: 13, Type: jdcal.Julian}
	nonLeapDaysPerMonth := nonLeapCyr.DaysPerMonth()
	output(`/*
NonLeapMonthProgression holds per month the ordinal of a day,
January 1st being 0, December 31st being 364 (365 days in a year, so ordinal number 364).
*/
`)
	output("var NonLeapMonthProgression = MonthProgression{\n")
	count = 0
	for m := time.January; m <= time.December; m++ {
		output("time.%v: {\n", m)
		output("0, // filler\n")
		for d := 1; d <= nonLeapDaysPerMonth[m]; d++ {
			output("%d, ", count)
			count++
			nonLeapYearDays = count // updated every time, don't care, it's code generation
		}
		output("\n},\n")
	}
	output("}\n\n")

	output("var progressionTableStart Year = %d\n", startYear)
	output("var progressionTableEnd Year = %d\n", endYear)

	output(`/*
YearProgression holds the number of days since jdcal.StartProgressionYear.
The first int in each entry is for jdcal.Gregorian, the second for jdcal.Julian.
*/
`)
	output("var YearProgression = map[Year][2]Ordinal{\n")
	gregorianProgression := 0
	julianProgression := startOffset
	for yr := startYear; yr <= endYear; yr++ {
		output("%d: {%d, %d}, // difference on Jan 1st: %d days\n",
			yr, gregorianProgression, julianProgression, julianProgression-gregorianProgression)

		cyr := jdcal.CalendarYear{Year: jdcal.Year(yr), Type: jdcal.Gregorian}
		if cyr.IsLeap() {
			gregorianProgression += leapYearDays
		} else {
			gregorianProgression += nonLeapYearDays
		}

		cyr.Type = jdcal.Julian
		if cyr.IsLeap() {
			julianProgression += leapYearDays
		} else {
			julianProgression += nonLeapYearDays
		}
	}
	output("}\n")

	check(os.WriteFile(os.Args[1], []byte(strings.Join(outputLines, "")), 0644))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func output(f string, args ...any) {
	outputLines = append(outputLines, fmt.Sprintf(f, args...))
}

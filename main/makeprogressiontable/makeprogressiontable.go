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

const (
	startYear   = -500 // generate from here
	endYear     = 2100 // until here
	startOffset = -5   // Gregorian is 5 days behind on startYear

	header = `package jdcal

// This is a generated file, don't edit. Look for a nearby Makefile.
	
import "time"

const (
	StartProgressionYear = %d
	EndProgressionYear   = %d

)
`
)

var outputLines = []string{}

func main() {
	if len(os.Args) != 2 {
		check(errors.New("usage: go run main/progressiontable/progressiontable.go $OUTPUTFILE"))
	}

	outputLines = append(outputLines, fmt.Sprintf(header, startYear, endYear))

	var leapYearDays, nonLeapYearDays int

	leapCyr := jdcal.CalendarYear{Year: 12, Type: jdcal.Julian}
	leapDaysPerMonth := leapCyr.DaysPerMonth()
	output(`/*
LeapMonthProgression holds per month the ordinal of a day,
January 1st being 1, December 31st being 366 (365 days in a year, plus February 29th).
*/
`)
	output("var LeapMonthProgression = map[time.Month][]int{\n")
	count := 0
	for m := time.January; m <= time.December; m++ {
		output("time.%v: {\n", m)
		output("0, // filler\n")
		for d := 1; d <= leapDaysPerMonth[m]; d++ {
			count++
			output("%d, ", count)
			leapYearDays = count // updated every time, don't care, it's code generation
		}
		output("\n},\n")
	}
	output("}\n\n")

	nonLeapCyr := jdcal.CalendarYear{Year: 13, Type: jdcal.Julian}
	nonLeapDaysPerMonth := nonLeapCyr.DaysPerMonth()
	output(`/*
NonLeapMonthProgression holds per month the ordinal of a day,
January 1st being 1, December 31st being 365.
*/
`)
	output("var NonLeapMonthProgression = map[time.Month][]int{\n")
	count = 0
	for m := time.January; m <= time.December; m++ {
		output("time.%v: {\n", m)
		output("0, // filler\n")
		for d := 1; d <= nonLeapDaysPerMonth[m]; d++ {
			count++
			output("%d, ", count)
			nonLeapYearDays = count // updated every time, don't care, it's code generation
		}
		output("\n},\n")
	}
	output("}\n\n")

	output(`/*
YearProgression holds the number of days since jdcal.StartProgressionYear.
The first int in each entry is for jdcal.Gregorian, the second for jdcal.Julian.
*/
`)
	output("var YearProgression = map[Year][2]int{\n")
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

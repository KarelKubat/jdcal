package timeline

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	daysFlag      = "days"
	repeatFlag    = "repeat"
	longUsage     = `
Prints a timeline of two calendars. Examples:

  jdcal timeline 1582/03/15              # timeline starts on March 15th 1582 on the Julian calendar
  jdcal timeline 1582/03/25 --gregorian  # reference date is taken as a Gregorian date
  jdcal timeline 1582/03/25 --days 60    # Julian calendar, over 60 days
  jdcal timeline 1582/03/15 --repeat     # repeat YYYY/MM in subsequent dates`
)

var Cmd = &cobra.Command{
	Use:   "timeline",
	Short: "Print a timeline of two calendars",
	Long:  longUsage,
	Args:  cobra.ExactArgs(1),
	Run:   runTimeline,
}

func init() {
	Cmd.Flags().Bool(gregorianFlag, false, "reference date is Gregorian")
	Cmd.Flags().Int(daysFlag, 30, "days to display")
	Cmd.Flags().Bool(repeatFlag, false, "repeat years and months in next lines")
}

func runTimeline(cmd *cobra.Command, args []string) {
	tp := jdcal.Julian
	gotG, err := cmd.Flags().GetBool(gregorianFlag)
	check(err)
	if gotG {
		tp = jdcal.Gregorian
	}
	dt, err := jdcal.NewFromString(args[0], tp)
	check(err)

	days, err := cmd.Flags().GetInt(daysFlag)
	check(err)

	repeat, err := cmd.Flags().GetBool(repeatFlag)
	check(err)

	if dt.Type == jdcal.Julian {
		fmt.Printf("%10s | %10s\n", "Julian", "Gregorian")
	} else {
		fmt.Printf("%10s | %10s\n", "Gregorian", "Julian")
	}
	fmt.Println("---------- | ----------")

	var lastDtYear, lastOtYear int
	var lastDtMonth, lastOtMonth time.Month
	for i := 0; i < days; i++ {
		ot, err := dt.Convert()
		check(err)
		printDate(dt, lastDtYear, lastDtMonth)
		fmt.Printf(" | ")
		printDate(ot, lastOtYear, lastOtMonth)
		fmt.Println()
		if !repeat {
			lastDtYear = dt.Year
			lastDtMonth = dt.Month
			lastOtYear = ot.Year
			lastOtMonth = ot.Month
		}
		// fmt.Println(dt, "is", ot)
		dt = dt.Advance()
	}
}

func printDate(d jdcal.Date, lastYear int, lastMonth time.Month) {
	if d.Year != lastYear {
		fmt.Printf("%4.4d/", d.Year)
	} else {
		fmt.Printf("     ")
	}
	if d.Month != lastMonth {
		fmt.Printf("%2.2d/", int(d.Month))
	} else {
		fmt.Printf("   ")
	}
	fmt.Printf("%2.2d", d.Day)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

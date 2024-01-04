package timeline

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	daysFlag      = "days"
	repeatFlag    = "repeat"
	zoneFlag      = "zone"
	longUsage     = `
Prints a timeline of two calendars. Examples:

  jdcal timeline 1582/10/01              # timeline starts onOctober 1st 1582 on the Julian calendar
  jdcal timeline 1582/10/11 --gregorian  # reference date is taken as a Gregorian date
  jdcal timeline 1582/10/01 --days 60    # Julian calendar, over 60 days
  jdcal timeline 1582/10/01 --zone spain # Show how Spain's calendar progressed
  jdcal timeline 1582/10/01 --repeat     # repeat YYYY/MM in subsequent dates`
)

var Cmd = &cobra.Command{
	Use:   "timeline",
	Short: "Print a timeline of two calendars",
	Long:  longUsage,
	Args:  cobra.ExactArgs(1),
	Run:   runTimeline,
}

func init() {
	Cmd.Flags().BoolP(gregorianFlag, strings.Split(gregorianFlag, "")[0], false,
		"reference date is Gregorian")
	Cmd.Flags().IntP(daysFlag, strings.Split(daysFlag, "")[0], 30, "days to display")
	Cmd.Flags().BoolP(repeatFlag, strings.Split(repeatFlag, "")[0], false,
		"repeat years and months in next lines")
	Cmd.Flags().StringP(zoneFlag, strings.Split(zoneFlag, "")[0], "", "show zone progress")
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

	var zones []jdcal.ZoneEntry
	zoneName, err := cmd.Flags().GetString(zoneFlag)
	check(err)
	if zoneName != "" {
		zones = jdcal.ZonesByName(zoneName)
		if len(zones) == 0 {
			check(fmt.Errorf("zone %q does not match anything, try `jdcal zones`", zoneName))
		}
		if len(zones) > 1 {
			check(fmt.Errorf("zone %q matches multiple zones, restrict the --zone name", zoneName))
		}
		fmt.Println(zones[0])
	}

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
		dtPrinted := printDate(dt, lastDtYear, lastDtMonth, zones)
		fmt.Printf(" | ")
		otPrinted := printDate(ot, lastOtYear, lastOtMonth, zones)
		fmt.Println()
		if !repeat {
			if dtPrinted {
				lastDtYear = dt.Year
				lastDtMonth = dt.Month
			}
			if otPrinted {
				lastOtYear = ot.Year
				lastOtMonth = ot.Month
			}
		}
		dt = dt.Advance()
	}
}

func printDate(d jdcal.Date, lastYear int, lastMonth time.Month, zones []jdcal.ZoneEntry) bool {
	if len(zones) == 1 {
		in, err := d.InZone(zones[0])
		check(err)
		if !in {
			fmt.Printf("          ")
			return false
		}
	}
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
	return true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
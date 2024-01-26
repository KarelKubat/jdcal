package holidays

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	firstHoliday  = jdcal.AshWednesday
	lastHoliday   = jdcal.Pentecost
	gregorianFlag = "gregorian"
	zoneFlag      = "zone"
	longUsage     = `
Prints the holidays for a given year. Examples:

jdcal holidays 1300                   # Holidays in 1300, Julian output format (default)
jdcal holidays 1999 --gregorian       # Holidays in 1999, Gregorian output format
jdcal holidays 1583 -z ausserrhoden   # Holidays in 1583, output format is whatever matches the zone`
)

var Cmd = &cobra.Command{
	Use:   "holidays",
	Short: "Print the holiday dates for a given year",
	Long:  longUsage,
	Args:  cobra.MinimumNArgs(1),
	Run:   runHolidays,
}

func init() {
	Cmd.Flags().BoolP(gregorianFlag, strings.Split(gregorianFlag, "")[0], false,
		"output format for dates is Gregorian, default: Julian")
	Cmd.Flags().StringP(zoneFlag, strings.Split(zoneFlag, "")[0], "",
		"derive output format from the zone")
}

func runHolidays(cmd *cobra.Command, args []string) {
	gotG, err := cmd.Flags().GetBool(gregorianFlag)
	check(err)

	var zone jdcal.ZoneEntry
	zoneName, err := cmd.Flags().GetString(zoneFlag)
	check(err)
	if zoneName != "" {
		zone, err = jdcal.SingleZone(zoneName)
		check(err)
	}

	if gotG && zoneName != "" {
		check(errors.New("--zone and --gregorian are mutually exclusive flags"))
	}

	for i, arg := range args {
		yr := atoi(arg)

		// Always fetch holidays relative to the Gregorian calendar. Convert to Julian later, if
		// that output format is requested, or if we're showing the date for a zone which at that time
		// is in Julian mode.
		cyr, err := jdcal.NewCalendarYear(jdcal.Year(yr), jdcal.Gregorian)
		check(err)

		if i > 0 {
			fmt.Println()
		}
		for h := firstHoliday; h <= lastHoliday; h++ {
			dt, err := cyr.HolidayDate(h)
			check(err)

			if !gotG {
				dt, err = dt.Convert()
				check(err)
			}
			if zoneName != "" {
				in, err := dt.InZone(zone)
				check(err)
				if !in {
					dt, err = dt.Convert()
					check(err)
				}
			}
			fmt.Printf("%-13s", h)
			if zoneName != "" {
				fmt.Printf(" in %s", zone.Name)
			}
			wd, err := dt.Weekday()
			check(err)
			fmt.Printf(" occurs on %-9s %v\n", wd, dt)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func atoi(s string) int {
	r, err := strconv.Atoi(s)
	check(err)
	return r
}

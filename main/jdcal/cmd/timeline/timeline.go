package timeline

import (
	"fmt"
	"log"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	daysFlag      = "days"
	longUsage     = `
Prints a timeline of two calendars. Examples:

  jdcal timeline 1582/03/15              # timeline starts on March 15th 1582 on the Julian calendar
  jdcal timeline 1582/03/25 --gregorian  # reference date is taken as a Gregorian date
  jdcal timeline 1582/03/25 --days 60    # Julian calendar, over 60 days`
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
	for i := 0; i < days; i++ {
		ot, err := dt.Convert()
		check(err)
		fmt.Println(dt, "is", ot)
		dt = dt.Advance()
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

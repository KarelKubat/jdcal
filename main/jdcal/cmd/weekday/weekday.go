package weekday

import (
	"fmt"
	"log"
	"strings"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	longUsage     = `
Prints the weekday for a given date. Examples:

jdcal weekday 1594/11/01      # Nov. 1st, 1594 on the Julian calendar is a Tuesday
jdcal weekday 1594/11/01 -g   # But it's a Friday on the Gregorian calendar`
)

var Cmd = &cobra.Command{
	Use:   "weekday",
	Short: "Print the weekday for a given date",
	Long:  longUsage,
	Args:  cobra.MinimumNArgs(1),
	Run:   runWeekday,
}

func init() {
	Cmd.Flags().BoolP(gregorianFlag, strings.Split(gregorianFlag, "")[0], false,
		"dates are Gregorian")
}

func runWeekday(cmd *cobra.Command, args []string) {
	tp := jdcal.Julian
	gotG, err := cmd.Flags().GetBool(gregorianFlag)
	check(err)
	if gotG {
		tp = jdcal.Gregorian
	}

	for _, arg := range args {
		dt, err := jdcal.NewDateFromString(arg, tp)
		check(err)
		wd, err := dt.Weekday()
		check(err)
		fmt.Println(dt, "is a", wd)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package convert

import (
	"fmt"
	"log"
	"strings"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	julianFlag    = "julian"
	longUsage     = `
Converts from Julian to Gregorian, or vv. Examples:

  jdcal convert 1582/03/15 --julian     # what is March 15th, 1582 as a Gregorian date
  jdcal convert 1582/03/25 --gregorian  # what is March 25th, 1582 as a Julian date`
)

var Cmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the given date(s) from Julian to Gregorian or vv.",
	Long:  longUsage,
	Args:  cobra.MinimumNArgs(1),
	Run:   runConvert,
}

func init() {
	Cmd.Flags().BoolP(julianFlag, strings.Split(julianFlag, "")[0], false,
		"convert Julian to Gregorian")
	Cmd.Flags().BoolP(gregorianFlag, strings.Split(gregorianFlag, "")[0], false,
		"convert Gregorian to Julian")
}

func runConvert(cmd *cobra.Command, args []string) {
	gotG, err := cmd.Flags().GetBool(gregorianFlag)
	check(err)
	gotJ, err := cmd.Flags().GetBool(julianFlag)
	check(err)
	for _, a := range args {
		if gotG {
			dt, err := jdcal.NewFromString(a, jdcal.Gregorian)
			check(err)
			ot, err := dt.Convert()
			check(err)
			fmt.Println(dt, "is", ot)
		}
		if gotJ {
			dt, err := jdcal.NewFromString(a, jdcal.Julian)
			check(err)
			ot, err := dt.Convert()
			check(err)
			fmt.Println(dt, "is", ot)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

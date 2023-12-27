package convert

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	gregorianFlag = "gregorian"
	julianFlag    = "julian"
)

var Cmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the given date(s) from Julian to Gregorian or vv.",
	Args:  cobra.MinimumNArgs(1),
	Run:   runConvert,
}

func init() {
	Cmd.Flags().Bool(julianFlag, false, "convert Julian to Gregorian")
	Cmd.Flags().Bool(gregorianFlag, false, "convert Gregorian to Julian")
}

func runConvert(cmd *cobra.Command, args []string) {
	for _, a := range args {
		gotG, err := cmd.Flags().GetBool(gregorianFlag)
		check(err)
		if gotG {
			dt := argToDate(a, jdcal.Gregorian)
			ot, err := dt.Convert()
			check(err)
			fmt.Println(dt, "is", ot)
		}
		gotJ, err := cmd.Flags().GetBool(julianFlag)
		check(err)
		if gotJ {
			dt := argToDate(a, jdcal.Julian)
			ot, err := dt.Convert()
			check(err)
			fmt.Println(dt, "is", ot)
		}
	}
}

func argToDate(arg string, tp jdcal.Type) jdcal.Date {
	parts := strings.Split(arg, "/")
	if len(parts) != 3 {
		check(fmt.Errorf("malformed date %q, want YYYY/MM/DD", arg))
	}
	dt, err := jdcal.New(atoi(arg, parts[0]), time.Month(atoi(arg, parts[1])),
		atoi(arg, parts[2]), tp)
	check(err)
	return dt
}

func atoi(fullArg, arg string) int {
	r, err := strconv.Atoi(arg)
	if err != nil {
		check(fmt.Errorf("part %q of %q is not a number: %v", arg, fullArg, err))
	}
	return r
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

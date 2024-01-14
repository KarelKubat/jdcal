package zones

import (
	"fmt"
	"log"
	"strings"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	zoneFlag  = "zone"
	shortFlag = "short"
	longUsage = `
Shows zones and the dates where calendars were switched. Examples:

  jdcal zones                     # show all zones
  jdcal zones --short             # short: only zone names
  jdcal zones --zone switzerland  # show zones that match, case-insensitive
`
)

var Cmd = &cobra.Command{
	Use:   "zones",
	Short: "List zones and conversion dates",
	Long:  longUsage,
	Args:  cobra.MaximumNArgs(0),
	Run:   runZones,
}

func init() {
	Cmd.Flags().StringP(zoneFlag, strings.Split(zoneFlag, "")[0], "",
		"restrict to zones matching this substring")
	Cmd.Flags().BoolP(shortFlag, strings.Split(shortFlag, "")[0], false,
		"show zone names only")
}

func runZones(cmd *cobra.Command, args []string) {
	match, err := cmd.Flags().GetString(zoneFlag)
	check(err)

	short, err := cmd.Flags().GetBool(shortFlag)
	check(err)

	if match != "" {
		for i, z := range jdcal.ZonesByName(match) {
			if short {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s", z.Name)
			} else {
				fmt.Println(z)
			}
		}
		if short {
			fmt.Println()
		}
		return
	}
	for i, z := range jdcal.ZonesTable {
		if short {
			if i > 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%s ", z.Name)
		} else {
			fmt.Println(z)
		}
	}
	if short {
		fmt.Println()
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

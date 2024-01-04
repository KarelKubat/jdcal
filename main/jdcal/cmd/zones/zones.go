package zones

import (
	"fmt"
	"log"
	"strings"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	matchFlag = "match"
	longUsage = `
Shows zones and the dates where calendars were switched. Examples:

  jdcal zones                      # show all zones
  jdcal zones --match switzerland  # show zones that match, case-insensitive
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
	Cmd.Flags().StringP(matchFlag, strings.Split(matchFlag, "")[0], "", "restrict to zones matching this substring")
}

func runZones(cmd *cobra.Command, args []string) {
	match, err := cmd.Flags().GetString(matchFlag)
	check(err)
	if match != "" {
		for _, z := range jdcal.ZonesByName(match) {
			fmt.Println(z)
		}
		return
	}
	for _, z := range jdcal.ZonesTable {
		fmt.Println(z)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

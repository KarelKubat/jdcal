package zones

import (
	"fmt"
	"log"

	"github.com/KarelKubat/jdcal"
	"github.com/spf13/cobra"
)

const (
	matchFlag = "match"
)

var Cmd = &cobra.Command{
	Use:   "zones",
	Short: "List zones and conversion dates",
	Args:  cobra.MaximumNArgs(0),
	Run:   runZones,
}

func init() {
	Cmd.Flags().String(matchFlag, "", "restrict to zones matching this substring")
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

package main

import (
	"github.com/KarelKubat/jdcal/main/jdcal/cmd"
	"github.com/KarelKubat/jdcal/main/jdcal/cmd/convert"
	"github.com/KarelKubat/jdcal/main/jdcal/cmd/timeline"
	"github.com/KarelKubat/jdcal/main/jdcal/cmd/zones"
)

func main() {
	cmd.Root.AddCommand(convert.Cmd)
	cmd.Root.AddCommand(timeline.Cmd)
	cmd.Root.AddCommand(zones.Cmd)
	cmd.Root.Execute()
}

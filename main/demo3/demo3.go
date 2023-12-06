package main

import (
	"fmt"

	"github.com/KarelKubat/jdcal"
)

func main() {
	for _, e := range jdcal.ZonesByName("netherlands") {
		fmt.Println(e)
	}

	// Output (actual string representations may differ):
	//   Belgium (Southern Netherlands)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1582/12/20
	//   Netherlands (Brabant)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1582/12/14
	//   Netherlands (Drenthe)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1701/04/30
	//   Netherlands (Frisia)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1701/12/31
	//   Netherlands (Gelderland)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1700/06/12
	//   Netherlands (Groningen City)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1583/01/01
	// 	Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	// 	Switched to   the Gregorian calendar   on   Julian 1700/12/31
	//   Netherlands (Holland)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1583/01/01
	//   Netherlands (Utrecht, Overijssel)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1700/11/30
	//   Netherlands (Zeeland, States General)
	// 	Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 	Switched to   the Gregorian calendar   on   Julian 1582/12/14
}

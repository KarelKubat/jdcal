# jdcal

Go library and CLI utility to convert to-and-fro between Julian and Gregorian calendars.

<!-- toc -->
- [Short CLI synopsis](#short-cli-synopsis)
- [Short library synopsis](#short-library-synopsis)
  - [Conversions](#conversions)
  - [Honoring leap years](#honoring-leap-years)
  - [Zones](#zones)
  - [More documentation](#more-documentation)
<!-- /toc -->

## Short CLI synopsis

To install `jdcal` as a CLI tool, run `go install main/jdcal/jdcal.go`. After this you can use the utility to convert dates:

```sh
# Take October 5th 1582 as a Julian date and convert
jdcal convert --julian 1582/10/5
Julian 1582/10/05 is Gregorian 1582/10/15

# Take October 10th 1582 as a Gregorian and convert back
jdcal convert --gregorian 1582/10/15
Gregorian 1582/10/15 is Julian 1582/10/05
```

Start `jdcal` without any arguments to see the usage information. The CLI tool performs more than just conversions.

## Short library synopsis

### Conversions

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	// October 5th (Julian) was the Papal announcement to skip 10 days.
	// The new date would be October 15 (Gregorian).
	jd0, err := jdcal.New(1582, time.October, 5, jdcal.Julian)
	check(err)

	// to Gregorian
	gd, err := jd0.Convert()
	check(err)
	fmt.Println("From Julian to Gregorian:", jd0, "is", gd)

	// back to Julian
	jd1, err := gd.Convert()
	check(err)
	fmt.Println("And back again:", gd, "is", jd1)

	// Output:
	// From Julian to Gregorian: Julian 1582/10/05 is Gregorian 1582/10/15
	// And back again: Gregorian 1582/10/15 is Julian 1582/10/05
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
```

### Honoring leap years

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

func main() {
	jd, err := jdcal.New(300, time.February, 26, jdcal.Julian)
	check(err)

	for i := 0; i < 10; i++ {
		gd, err := jd.Convert()
		check(err)
		fmt.Println(jd, "is", gd)
		jd = jd.Advance()
	}

	// Output:
	// Julian 0300/02/26 is Gregorian 0300/02/26
	// Julian 0300/02/27 is Gregorian 0300/02/27
	// Julian 0300/02/28 is Gregorian 0300/02/28
	// Julian 0300/02/29 is Gregorian 0300/03/01 <-- Note: leap year on Julian.
	// Julian 0300/03/01 is Gregorian 0300/03/02     Now there's a 1 day difference.
	// Julian 0300/03/02 is Gregorian 0300/03/03
	// Julian 0300/03/03 is Gregorian 0300/03/04
	// Julian 0300/03/04 is Gregorian 0300/03/05
	// Julian 0300/03/05 is Gregorian 0300/03/06
	// Julian 0300/03/06 is Gregorian 0300/03/07
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
```

### Zones

The package knows about zones, and on which dates these zones switched. Some zones even first switched from Julian to Gregorian, but then back again to Julian, and then forward.

The below code would display zone information in a human readable way, tough a `jdcal.ZoneEntry` holds this as a `struct` that can be programmatically examined.

```go
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
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1582/12/20
	//   Netherlands (Brabant)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1582/12/14
	//   Netherlands (Drenthe)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1701/04/30
	//   Netherlands (Frisia)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1701/12/31
	//   Netherlands (Gelderland)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1700/06/12
	//   Netherlands (Groningen City)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1583/01/01
	// 		Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	// 		Switched to   the Gregorian calendar   on   Julian 1700/12/31
	//   Netherlands (Holland)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1583/01/01
	//   Netherlands (Utrecht, Overijssel)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1700/11/30
	//   Netherlands (Zeeland, States General)
	// 		Started using the Julian    calendar   on   Gregorian -0500/02/28
	// 		Switched to   the Gregorian calendar   on   Julian 1582/12/14
}
```

### More documentation

For more please see the `doc.go`-generated docs.

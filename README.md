# jdcal

Go library and CLI utility to convert to-and-fro between Julian and Gregorian calendars.

<!-- toc -->
- [Short CLI synopsis](#short-cli-synopsis)
  - [Conversions between calendars](#conversions-between-calendars)
  - [When did the world switch calendars](#when-did-the-world-switch-calendars)
  - [Visualization of a timeline](#visualization-of-a-timeline)
- [Short library synopsis](#short-library-synopsis)
  - [Conversions](#conversions)
  - [Honoring leap years](#honoring-leap-years)
  - [Zones](#zones)
  - [Testing whether a date exists in a zone](#testing-whether-a-date-exists-in-a-zone)
  - [More documentation](#more-documentation)
<!-- /toc -->

## Short CLI synopsis

To install `jdcal` as a CLI tool, run `go install main/jdcal/jdcal.go`. After this you can use the utility. Start `jdcal` without any arguments to see the usage information.

### Conversions between calendars

```sh
# Take October 5th 1582 as a Julian date and convert
jdcal convert --julian 1582/10/05  # or use -j as a short flag
Julian 1582/10/05 is Gregorian 1582/10/15

# Take October 10th 1582 as a Gregorian and convert back
jdcal convert --gregorian 1582/10/15  # or use -g as a short flag
Gregorian 1582/10/15 is Julian 1582/10/05
```

### When did the world switch calendars

```sh
# Switch over dates for zones matching "america"
jdcal zones --match america
United States of America (British Empire)
  Started using the Julian    calendar   on   Gregorian -0500/02/28
  Switched to   the Gregorian calendar   on   Julian 1752/09/02
United States of America (Russion Empire: Alaska)
  Started using the Julian    calendar   on   Gregorian -0500/02/28
  Switched to   the Gregorian calendar   on   Julian 1867/10/06

# All known zones
jdcal zones
  # Shows a long list
```

### Visualization of a timeline

Spain switched over to the Gregorian calendar on October 4th 1582 (try it with `jdcal zones --match spain`). A visualization:

```sh
# Show the timeline around 1582/10/04
jdcal timeline 1582/10/01 --days 10
    Julian   |    Gregorian
-----------+---+-----------
1582/10/01 | M | 1582/10/11
        02 | T |         12
        03 | W |         13
        04 | T |         14
        05 | F |         15
        06 | S |         16
        07 | S |         17
        08 | M |         18
        09 | T |         19
        10 | W |         20
```

So the Spanish people went to sleep on October 4th and woke up on the 15th. That year, valid October dates were 1, 2, 3, 4, 15, 16, 17, etc.. Dates like October 10th don't exist in that zone.

One can argue whether October 14th is a valid date as it is the Gregorian version of the switch-over date October 4th. `jdcal` assumes that it is.

The output can be further trimmed by using the flag `--zone`:

```sh
# Show the timeline around 1582/10/04 for zone Spain, which switched over to
# the Gregorian calendar on 1582/10/04.
jdcal timeline 1582/10/01 --days 10 --zone spain
    Julian   |    Gregorian
-----------+---+-----------
1582/10/01 | M |
        02 | T |
        03 | W |
        04 | T | 1582/10/14
           | F |         15
           | S |         16
           | S |         17
           | M |         18
           | T |         19
           | W |         20
```

Zones that switch back from Gregorian to Julian offer interesting timelines. E.g., "Groningen City" temporarily switched back to Julian after being re-conquered by Protestants. This happened on (the Gregorian date) 1594/11/10, which became (Julian) 1594/10/31. (Again, one may argue whether the switch-over date Julian 1594/10/31 is valid. `jdcal` assumes it is.) A visualization:

```sh
jdcal zones -m 'groningen city'
Netherlands (Groningen City)
  Started using the Julian    calendar   on   Gregorian -0500/02/28
  Switched to   the Gregorian calendar   on   Julian 1583/01/01
  Switched to   the Julian    calendar   on   Gregorian 1594/11/10
  Switched to   the Gregorian calendar   on   Julian 1700/12/31

jdcal timeline 1594/11/01 --days  20 --zone 'groningen city' --gregorian
 Gregorian |   |     Julian
-----------+---+-----------
1594/11/01 | T |
        02 | W |
        03 | T |
        04 | F |
        05 | S |
        06 | S |
        07 | M |
        08 | T |
        09 | W |
        10 | T | 1594/10/31
           | F |      11/01
           | S |         02
           | S |         03
           | M |         04
           | T |         05
           | W |         06
           | T |         07
           | F |         08
           | S |         09
           | S |         10
```

In this case a date like 1594/11/08 exists twice. Also;

- The date 1594/10/30 must be Gregorian;
- Dates 1594/10/31 until 1594/11/10 exist on both calendars;
- The date 1594/11/11 must be Julian.

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
	wd, err := gd.Weekday()
	check(err)
	fmt.Println("From Julian to Gregorian:", jd0, "is", gd, "and it's a", wd)

	// back to Julian
	jd1, err := gd.Convert()
	check(err)
	wd, err = jd1.Weekday()
	check(err)
	fmt.Println("And back again:", gd, "is", jd1, "and it's a", wd)

	// Output:
	// 	From Julian to Gregorian: Julian 1582/10/05 is Gregorian 1582/10/15 and it's a Friday
	// 	And back again: Gregorian 1582/10/15 is Julian 1582/10/05 and it's a Friday
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

### Testing whether a date exists in a zone

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

const (
	zoneName = "Netherlands (Groningen City)"
)

func main() {
	zones := jdcal.ZonesByName(zoneName)
	if len(zones) != 1 {
		check(fmt.Errorf("name %q does not select one single zone", zoneName))
	}

	fmt.Println(zones[0])
	// Netherlands (Groningen City)
	//   Started using the Julian    calendar   on   Gregorian -0500/02/28
	//   Switched to   the Gregorian calendar   on   Julian 1583/01/01
	//   Switched to   the Julian    calendar   on   Gregorian 1594/11/10
	//   Switched to   the Gregorian calendar   on   Julian 1700/12/31

	// Some dates that lie in between the cutovers.
	for _, year := range []int{1580, 1590, 1600, 1800} {
		test(year, time.January, 1, zones[0])
	}

	// Just around the exact cutover dates.
	for _, date := range []struct {
		year  int
		month time.Month
		day   int
	}{
		// Around the first switch from Julian into the Gregorian calendar
		{year: 1582, month: time.December, day: 31},
		{year: 1583, month: time.January, day: 1},
		{year: 1583, month: time.January, day: 2},

		// Around the second switch from Gregorian back to Julian
		{year: 1594, month: time.November, day: 9},
		{year: 1594, month: time.November, day: 10},
		{year: 1594, month: time.November, day: 11},

		// Around the third switch back to Gregorian
		{year: 1700, month: time.December, day: 30},
		{year: 1700, month: time.December, day: 31},
		{year: 1701, month: time.January, day: 1},
	} {
		test(date.year, date.month, date.day, zones[0])
	}
}

func test(year int, month time.Month, day int, z jdcal.ZoneEntry) {
	d, err := jdcal.New(year, month, day, jdcal.Julian)
	check(err)
	jdInZone, err := d.InZone(z)
	check(err)

	d, err = jdcal.New(year, month, day, jdcal.Gregorian)
	check(err)
	gdInZone, err := d.InZone(z)
	check(err)

	fmt.Printf("%4.4d/%2.2d/%2.2d ", year, int(month), day)
	switch {
	case jdInZone && gdInZone:
		fmt.Println("can be both a Julian and a Gregorian date")
	case !jdInZone && !gdInZone:
		fmt.Println("is neither a Julian nor a Gregorian date")
	case jdInZone:
		fmt.Println("is a Julian date")
	default:
		fmt.Println("is a Gregorian date")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
```

### More documentation

For more please see the `doc.go`-generated docs.

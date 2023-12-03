# jdcal

Go utility to convert to-and-fro between Julian and Gregorian calendars.

<!-- toc -->
- [Shortest synopsis](#shortest-synopsis)
<!-- /toc -->

## Shortest synopsis

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
	jd0, err := jdcal.New(jdcal.Date{
		Year:  1582,
		Month: time.October,
		Day:   5,
		Type:  jdcal.Julian,
	})
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

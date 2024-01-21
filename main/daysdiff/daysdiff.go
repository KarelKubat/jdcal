package main

import (
	"fmt"
	"log"
	"time"

	"github.com/KarelKubat/jdcal"
)

const (
	referenceMonth = time.March
	referenceDay   = 21
	endYear        = 2100
)

/*
Prints a table suitable for GnuPlot of:

- Year, being only the century year (so 1200, 1300 etc.)

- # of days on the Julian calendar since start of epoch

- # of days on the Gregorian calendar since start of epoch

- Difference between to two, i.e., by how many days was the Julian calendar wrong
on this date.

Example:

-500, 16, 21, -5
-400, 36541, 36546, -5
-300, 73066, 73070, -4
-200, 109591, 109594, -3
-100, 146116, 146118, -2
100, 218800, 218801, -1
200, 255325, 255325, 0
300, 291850, 291849, 1
400, 328375, 328374, 1
500, 364900, 364898, 2
600, 401425, 401422, 3
700, 437950, 437946, 4
800, 474475, 474471, 4
900, 511000, 510995, 5
1000, 547525, 547519, 6
1100, 584050, 584043, 7
1200, 620575, 620568, 7
1300, 657100, 657092, 8
1400, 693625, 693616, 9
1500, 730150, 730140, 10
1600, 766675, 766665, 10
1700, 803200, 803189, 11
1800, 839725, 839713, 12
1900, 876250, 876237, 13
2000, 912775, 912762, 13
2100, 949300, 949286, 14
*/
func main() {
	for yr := jdcal.Year(-400); yr < 2100; yr += 100 {
		if yr == 0 {
			continue
		}
		jd, err := jdcal.NewDate(yr, referenceMonth, referenceDay, jdcal.Julian)
		check(err)
		gd, err := jdcal.NewDate(yr, referenceMonth, referenceDay, jdcal.Gregorian)
		check(err)

		jdOrd := jd.Ordinal()
		gdOrd := gd.Ordinal()

		fmt.Printf("%d, %d, %d, %d\n", yr, jdOrd, gdOrd, jdOrd-gdOrd)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

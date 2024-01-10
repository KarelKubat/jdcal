package main

import (
	"fmt"
	"slices"
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

- Reference date, being only the century year (so 1200, 1300 etc.)

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
	jd := jdcal.First(jdcal.Julian)
	gd := jdcal.First(jdcal.Gregorian)
	day := 0

	jdMark := map[int]int{}
	gdMark := map[int]int{}
	marks := []int{}
	endReferencesSeen := 0

	for {
		if isReference(jd.Year, jd.Month, jd.Day) {
			jdMark[jd.Year] = day
			if !slices.Contains(marks, jd.Year) {
				marks = append(marks, jd.Year)
			}
			if jd.Year == endYear {
				endReferencesSeen++
			}
		}
		if isReference(gd.Year, gd.Month, gd.Day) {
			gdMark[gd.Year] = day
			if !slices.Contains(marks, gd.Year) {
				marks = append(marks, gd.Year)
			}
			if gd.Year == endYear {
				endReferencesSeen++
			}
		}

		if endReferencesSeen == 2 {
			break
		}

		jd = jd.Advance()
		gd = gd.Advance()
		day++
	}

	for _, m := range marks {
		fmt.Printf("%d, %d, %d, %d\n", m, jdMark[m], gdMark[m], jdMark[m]-gdMark[m])
	}
}

func isReference(y int, m time.Month, d int) bool {
	return y%100 == 0 && m == referenceMonth && d == referenceDay
}

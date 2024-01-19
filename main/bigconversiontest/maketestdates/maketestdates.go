package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/KarelKubat/jdcal"
)

const (
	header = `package testdates

import (
	"time"
)

type Stamp struct {
	Year  int
	Month time.Month
	Day   int
}

type TwoDates struct {
	J, G Stamp
}

var TestDates = [...]TwoDates{
	`
	footer = `}
`
)

func main() {
	if len(os.Args) != 2 {
		check(errors.New("usage: main/maketestdates/maketestdates.go OUTPUTFILE"))
	}

	if _, err := os.Stat(os.Args[1]); err == nil {
		check(fmt.Errorf("%q already exists, won't overwrite", os.Args[1]))
	}

	f, err := os.Create(os.Args[1])
	check(err)
	defer f.Close()

	f.Write([]byte(header))

	jd := jdcal.First(jdcal.Julian)
	jd.Year++
	var prevYear jdcal.Year = -999999
	last := jdcal.Last(jdcal.Julian)
	last.Year--
	for {
		if jd.Year != prevYear {
			prevYear = jd.Year
			fmt.Printf("%v ", prevYear)
		}
		bf, err := jd.AfterOrEqual(last)
		check(err)
		if bf {
			break
		}

		gd, err := jd.Convert()
		check(err)

		f.Write([]byte(fmt.Sprintf(`
			{ 
				J: Stamp{Year: %d, Month: time.%s, Day: %d},
				G: Stamp{Year: %d, Month: time.%s, Day: %d},
			},
			`, jd.Year, jd.Month, jd.Day, gd.Year, gd.Month, gd.Day)))

		// Skip 3 days to keep the set at a reasonable size. Github doesn't like big files.
		jd = jd.Forward()
		jd = jd.Forward()
		jd = jd.Forward()

		/* Temp */
		// if jd.Year == -300 {
		// 	break
		// }
	}
	fmt.Println()
	f.Write([]byte(footer))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

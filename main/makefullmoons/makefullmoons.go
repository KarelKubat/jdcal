package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KarelKubat/jdcal"
)

const (
	urlFormat = "http://astropixels.com/ephemeris/phasescat/phases%.4d.html"
	startYear = -499
	endYear   = 2100

	header = `package jdcal

import "time"

/* 
FullMoons maps years to a month and day (MD) representing the first full moon beyond March 21st. The date points are relative to the Georgian calendar. Example:

  md, ok := jdcal.FullMoons[1600]
  if !ok {
	log.Fatal("no full moon information for year 1600")
  }
  fmt.Println("first full moon beyond March 21st in the year 1600 is on", md)  // 03/29

The data were obtained from https://astropixels.com/ephemeris/phasescat/phasescat.html
Moon Phases Table courtesy of Fred Espenak, www.Astropixels.com.

The data here differ from www.Astropixels.com in that:

- Years are stated as BC or AD - so -3, -2, -1, 1, 2, 3, etc.. There is no year zero.
Astropixels.com knows a year zero and has  -2, -1, 0, 1, 2, 3.

- All dates are relative to the Gregorian calendar. Astropixels.com uses Julian for pre-1582, this is converted to Gregorian for clarity.
*/
var FullMoons = map[Year]MD{
`

	trailer = `
}
`
)

func main() {
	if len(os.Args) != 2 {
		check(errors.New("usage: go run main/makefullmoons/makefullmoons.go OUTPUTFILE"))
	}

	out := header
	for y := startYear; y <= endYear; y += 100 {
		url := fmt.Sprintf(urlFormat, y)
		out += processURL(url)
	}
	out += trailer
	check(os.WriteFile(os.Args[1], []byte(out), 0644))
}

func processURL(url string) string {
	fmt.Printf("\nFetching %q...\n", url)
	resp, err := http.Get(url)
	check(err)
	if resp.StatusCode != http.StatusOK {
		check(fmt.Errorf("got code %v, need %v", resp.StatusCode, http.StatusOK))
	}
	body, err := io.ReadAll(resp.Body)
	check(err)
	lines := extractTables(string(body))

	return fmt.Sprintf("\n// %s\n%s", url, extractDates(lines))
}

/*
The blocks  we want have information between <pre> and </pre>. Inside the blocks we may see </div> and <br/>
*/
func extractTables(body string) (lines []string) {
	active := false
	for _, line := range strings.Split(body, "\n") {
		if strings.Contains(line, "<pre>") {
			active = true
		}
		if strings.Contains(line, "</pre>") || strings.Contains(line, "</div>") {
			active = false
		}
		if !active {
			continue
		}
		line = strings.ReplaceAll(line, "<br/>", "")
		line = strings.ReplaceAll(line, "<pre>", "")
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

/*
Relevant blocks are in the format:

	Year      New Moon        First Quarter       Full Moon       Last Quarter
	-0401                                       Jan  3  07:08     Jan 11  11:47
			Jan 18  08:20 T   Jan 25  01:17     Feb  2  02:13 p   Feb 10  01:43
			Feb 16  18:08     Feb 23  15:24     Mar  3  20:01     Mar 11  11:46
			Mar 18  03:06     Mar 25  07:24     Apr  2  11:14     Apr  9  18:46
			Apr 16  11:50     Apr 24  00:41     May  1  23:35     May  8  23:50
			May 15  21:18     May 23  18:13     May 31  09:39     Jun  7  04:18
			Jun 14  08:31     Jun 22  11:02     Jun 29  18:18     Jul  6  09:36
			Jul 13  22:08 A   Jul 22  02:28     Jul 29  02:26 p   Aug  4  17:12
			Aug 12  14:03     Aug 20  16:14     Aug 27  10:46     Sep  3  04:23
			Sep 11  07:27     Sep 19  04:20     Sep 25  19:55     Oct  2  19:55
			Oct 11  01:09     Oct 18  14:47     Oct 25  06:30     Nov  1  15:27
			Nov  9  18:08     Nov 16  23:47     Nov 23  19:06     Dec  1  13:19
			Dec  9  09:39     Dec 16  07:54     Dec 23  09:56     Dec 31  11:10

Negative years are one off; this is numbered as -2, -1, 0, 1, 2 etc. instead of BC / AD.

Full moon dates are at pos 44, "Mar 21" in the below example

	          1         2         3         4         5         6         7
	01234567890123456789012345678901234567890123456789012345678901234567890123456789
	        Mar  6  19:33     Mar 13  11:54     Mar 21  12:51     Mar 29  11:32

We only need to emit one value: the full moon beyond March 21st.
*/
var yearEmitted = map[int]bool{}

func extractDates(lines []string) (out string) {
	for _, l := range lines {
		out += fmt.Sprintf("// %s\n", l)
	}
	var year int
	var expectYear bool
	for _, line := range lines {
		if expectYear {
			expectYear = false
			year = atoi(line)
			if year <= 0 {
				year-- // Adjust for BC, we don't use year nr 0
			}
			fmt.Printf("%d ", year)
		}
		if strings.Contains(line, " Year") {
			expectYear = true
			continue
		}

		// Avoid more YMDs for this year.
		if yearEmitted[year] {
			continue
		}

		// Extract relevant data, emit if the entry is later than March 21st.
		if len(line) < 70 {
			check(fmt.Errorf("unexpected line %q", line))
		}
		monthString := line[44:47]
		dayString := line[48:50]
		if monthString == "   " {
			continue
		}
		m := extractMonth(monthString)
		d := atoi(dayString)

		var dt jdcal.Date
		var err error

		// AstroPixels.com has dates before 1582 as Julians. Convert to Gregorian for unity.
		if year < 1582 {
			dt, err = jdcal.New(jdcal.Year(year), m, d, jdcal.Julian)
			if err != nil && strings.Contains(err.Error(), "is before the first convertible date") {
				continue
			}
			checkWithLine(err, line)
			dt, err = dt.Convert()
			check(err)
		} else {
			dt, err = jdcal.New(jdcal.Year(year), m, d, jdcal.Gregorian)
			if err != nil && strings.Contains(err.Error(), "is after the last convertible date") {
				continue
			}
			checkWithLine(err, line)
		}
		refDate := jdcal.Date{Year: jdcal.Year(year), Month: 3, Day: 21, Type: jdcal.Gregorian}
		af, err := dt.After(refDate)
		check(err)

		if af {
			out += fmt.Sprintf("%d: {time.%v, %d}, // %s/%s from %s\n",
				year, dt.Month, dt.Day, monthString, dayString, strings.TrimSpace(line))
			yearEmitted[year] = true
		}
	}

	fmt.Println()
	return out
}

func extractMonth(s string) time.Month {
	nameMap := map[string]time.Month{
		"Jan": time.January, "Feb": time.February, "Mar": time.March, "Apr": time.April,
		"May": time.May, "Jun": time.June, "Jul": time.July, "Aug": time.August,
		"Sep": time.September, "Oct": time.October, "Nov": time.November, "Dec": time.December,
	}
	for k, v := range nameMap {
		if strings.HasPrefix(s, k) {
			return v
		}
	}
	check(fmt.Errorf("cannot get month from substring %q", s))
	return time.January
}

func atoi(line string) (nr int) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	nr, err := strconv.Atoi(parts[0])
	if err != nil {
		check(fmt.Errorf("failed to extract year from line %q: %v", line, err))
	}
	return nr
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkWithLine(err error, line string) {
	if err != nil {
		log.Fatalf("at line %q: %v", line, err)
	}
}

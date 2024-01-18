package jdcal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// StringToYMD is a simple string parser to convert a date in the format "YYYY/MM/DD" to a separate year, month and day. The separator is a slash, to allow for easier negative years (as in "-25/02/28"). An error occurs when the parts in the argument cannot be converted; the validity of the date is not checked (so "2020/12/100" would pass).
func StringToYMD(arg string) (ymd YMD, err error) {
	parts := strings.Split(arg, "/")
	if len(parts) != 3 {
		return YMD{}, fmt.Errorf("malformed date %q, want YYYY/MM/DD", arg)
	}
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return YMD{}, fmt.Errorf("year part %q of %q is not a number: %v", parts[0], arg, err)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return YMD{}, fmt.Errorf("month part %q of %q is not a number: %v", parts[1], arg, err)
	}
	month := time.Month(m)
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return YMD{}, fmt.Errorf("day part %q of %q is not a number: %v", parts[2], arg, err)
	}

	return YMD{Year: Year(year), Month: month, Day: day}, nil
}

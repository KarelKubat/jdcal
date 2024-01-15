package jdcal

import (
	"testing"
)

func TestHolidayString(t *testing.T) {
	// Iterate over used holiday enums, let it crash if holiday_string.go isn't complete.
	for h := firstUnusedHoliday + 1; h < lastUnusedHoliday; h++ {
		_ = h.String()
	}
}

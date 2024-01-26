package jdcal

import (
	"testing"
)

func TestAlgorithmString(t *testing.T) {
	for a := firstUnusedAlgorithm; a <= lastUnusedAlgorithm; a++ {
		_ = a.String() // let it crash if strings don't match the enum
	}
}

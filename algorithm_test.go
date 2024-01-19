package jdcal

import (
	"testing"
)

func TestAlgorithmSwitching(t *testing.T) {
	want := func(call string, wantAlgo int) {
		if algorithm != wantAlgo {
			t.Errorf("after %v(): algorithm = %v, want %v", call, algorithm, wantAlgo)
		}
	}

	ConvertByLookup()
	want("ConvertByLookup", algorithmLookupTable)
	ConvertByProgression()
	want("ConvertByProgression", algorithmProgression)
}

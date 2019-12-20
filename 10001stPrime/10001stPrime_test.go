package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestFindNPrimes(t *testing.T) {
	testCases := []struct {
		primeCount int
		want       []int64
	}{
		{0, []int64{}},
		{1, []int64{2}},
		{2, []int64{2, 3}},
		{6, []int64{2, 3, 5, 7, 11, 13}}, //the provided example
	}

	for _, c := range testCases {
		got := findNPrimes(c.primeCount)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("findNPrimes(%v) == %v want %v", c.primeCount, got, c.want)
		}
	}
}

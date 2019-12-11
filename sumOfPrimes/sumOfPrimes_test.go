package main

import (
	"testing"
)

func TestFactorNumber(t *testing.T) {
	testCases := []struct {
		max  int64
		want int64
	}{
		{0, 0},
		{1, 0},
		{2, 2},
		{10, 17}, //provied example
	}

	for _, c := range testCases {
		got := getSumOfPrimes(c.max)
		if got != c.want {
			t.Errorf("getSumOfPrimes(%v) == %v, want %v", c.max, got, c.want)
		}
	}
}

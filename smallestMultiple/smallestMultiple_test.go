package main

import (
	"testing"
)

func TestFindSmallestMultiple(t *testing.T) {
	testCases := []struct {
		min  int64
		max  int64
		want int64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 6},
		{1, 10, 2520}, //provided example
	}

	for _, c := range testCases {
		got := findSmallestMultiple(c.min, c.max)
		if got != c.want {
			t.Errorf("findSmallestMultiple(genFactors(%v, %v)) == %v want %v", c.min, c.max, got, c.want)
		}
	}
}

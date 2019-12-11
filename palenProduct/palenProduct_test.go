package main

import (
	"testing"
)

func TestFindLargestProduct(t *testing.T) {
	testCases := []struct {
		low  int64
		high int64
		want int64
	}{
		{1, 9, 9},
		{10, 99, 9009}, //the provided example
	}

	for _, c := range testCases {
		_, _, got := findLargestProduct(c.low, c.high)
		if got != c.want {
			t.Errorf("findEmirps(%v, %v) == %v want %v", c.low, c.high, got, c.want)
		}
	}
}

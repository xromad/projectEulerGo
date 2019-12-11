package main

import (
	"testing"
)

func TestExpInt64(t *testing.T) {
	testCases := []struct {
		x    int64
		n    int64
		want int64
	}{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 2, 1},
		{2, 2, 4},
		{2, 3, 8},
		{55, 2, 3025}, //provided example
	}

	for _, c := range testCases {
		got := expInt64(c.x, c.n)
		if got != c.want {
			t.Errorf("expInt64(%v, %v)) == %v want %v", c.x, c.n, got, c.want)
		}
	}
}

func TestGetSumOfSquares(t *testing.T) {
	testCases := []struct {
		min  int64
		max  int64
		want int64
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
		{1, 2, 5},
		{1, 3, 14},
		{1, 10, 385}, //provided example
	}

	for _, c := range testCases {
		got := getSumOfSquares(c.min, c.max)
		if got != c.want {
			t.Errorf("getSumOfSquares(%v, %v)) == %v want %v", c.min, c.max, got, c.want)
		}
	}
}

func TestGetSquareOfSum(t *testing.T) {
	testCases := []struct {
		min  int64
		max  int64
		want int64
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 2, 9},
		{1, 10, 3025}, //provided example
	}

	for _, c := range testCases {
		got := getSquareOfSum(c.min, c.max)
		if got != c.want {
			t.Errorf("getSumOfSquares(%v, %v)) == %v want %v", c.min, c.max, got, c.want)
		}
	}
}

// 1 - 10 square of the sum is 3025 − 385 = 2640.
func TestFindSumSquareDifference(t *testing.T) {
	testCases := []struct {
		min            int64
		max            int64
		wantSumSquares int64
		wantSquareSum  int64
		wantDifference int64
	}{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 2, 5, 9, 4},
		{1, 10, 385, 3025, 2640}, //provided example
	}

	for _, c := range testCases {
		gotSumSquares, gotSquareSum, gotDifference := findSumSquareDifference(c.min, c.max)
		if gotSumSquares != c.wantSumSquares || gotSquareSum != c.wantSquareSum || gotDifference != c.wantDifference {
			t.Errorf("findSumSquareDifference(%v, %v)) == %v, %v, %v want %v, %v, %v",
				c.min,
				c.max,
				gotSumSquares,
				gotSquareSum,
				gotDifference,
				c.wantSumSquares,
				c.wantSquareSum,
				c.wantDifference,
			)
		}
	}
}

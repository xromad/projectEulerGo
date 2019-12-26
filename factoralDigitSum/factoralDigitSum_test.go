package main

import (
	"testing"
)

func TestFactoralDigitSum(t *testing.T) {
	testCases := []struct {
		num  int64
		want int64
	}{
		//where
		{1, 1},
		{2, 2},
		{4, 6},
		{10, 27}, //provided example
	}
	for _, c := range testCases {
		//when
		got := factoralDigitSum(c.num)
		//then
		if got != c.want {
			t.Errorf("countSundays(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

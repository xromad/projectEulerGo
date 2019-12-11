package main

import (
	"testing"
)

//test genEquation
func TestGenEquation(t *testing.T) {
	testCases := []struct {
		start int
		end   int
		want  string
	}{
		{1, 10, "3+5+6+9=23"}, //provided example
		{1, 2, "0=0"},
		{0, 0, "Start must be >= 1"},
		{1, 101, "End must be <= 100"},
	}

	for _, c := range testCases {
		got := genEquation(c.start, c.end)
		if got != c.want {
			t.Errorf("genOutput(%d, %d) == %s, want %s", c.start, c.end, got, c.want)
		}
	}
}

func TestBuildString(t *testing.T) {
	testCases := []struct {
		intSlice []int
		want     string
	}{
		{[]int{3, 5, 6, 9}, "3+5+6+9=23"}, // provided example
		{[]int{}, "0=0"},
	}

	for _, c := range testCases {
		got := buildString(c.intSlice)
		if got != c.want {
			t.Errorf("sliceCompare(%d) == %s, want %s", c.intSlice, got, c.want)
		}
	}
}

func TestEqationDecider(t *testing.T) {
	testCases := []struct {
		num  int
		want bool
	}{
		{15, true},
		{5, true},
		{4, false},
		{3, true},
		{0, true},
	}

	for _, c := range testCases {
		got := eqDecider(c.num)
		if got != c.want {
			t.Errorf("getFizzBuzz(%d) == %t, want %t", c.num, got, c.want)
		}
	}
}

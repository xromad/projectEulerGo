package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestFactorNumber(t *testing.T) {
	testCases := []struct {
		number int64
		want   []int64
	}{
		{0, []int64{1}},
		{1, []int64{1}},
		{2, []int64{1, 2}},
		{3, []int64{1, 3}},
		{5, []int64{1, 5}},
		{7, []int64{1, 7}},
		{12, []int64{1, 2, 2, 3}},
		{15, []int64{1, 3, 5}},
		{13195, []int64{1, 5, 7, 13, 29}}, //the provided example
	}

	for _, c := range testCases {
		got := factorNumber(c.number)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("factorNumber(%v) == %v, want %v", c.number, got, c.want)
		}
	}
}

func TestGenEquation(t *testing.T) {
	testCases := []struct {
		number []int64
		want   string
	}{
		{[]int64{}, "0=0"},
		{[]int64{1}, "1=1"},
		{[]int64{1, 3, 5}, "1*3*5=15"},
	}

	for _, c := range testCases {
		got := genEquation(c.number)
		if got != c.want {
			t.Errorf("genEquation(%v) == %v, want %v", c.number, got, c.want)
		}
	}
}

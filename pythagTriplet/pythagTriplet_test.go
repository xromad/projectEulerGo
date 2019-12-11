package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestFindNextPythagTriplet(t *testing.T) {
	testCases := []struct {
		aIn  int64
		bIn  int64
		max  int64
		want []int64
	}{
		{1, 1, 100, []int64{3, 4, 5}}, //provided example
		{3, 5, 100, []int64{5, 12, 13}},
		{6, 9, 10, []int64{0, 0, 0}},
	}

	for _, c := range testCases {
		got := findNextPythagTriplet(c.aIn, c.bIn, c.max)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("findNextPythagTriplet(%v, %v, %v) == %v want %v", c.aIn, c.bIn, c.max, got, c.want)
		}
	}
}

func TestGetPythagTripletEqN(t *testing.T) {
	testCases := []struct {
		target int64
		want   []int64
	}{
		{5, []int64{0,0,0}},
		{12, []int64{3, 4, 5}}, //provided example
		{30, []int64{5, 12, 13}},
	}

	for _, c := range testCases {
		got := getPythagTripletEqN(c.target)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getPythagTripleEqN(%v) == %v want %v", c.target, got, c.want)
		}
	}
}

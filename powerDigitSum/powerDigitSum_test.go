package main

import (
	commonintutils "projectEuler/common/commonIntUtils"
	"testing"
)

func TestGetDigits(t *testing.T) {
	testCases := []struct {
		num  float64
		want []int
	}{
		//where
		{0, []int{0}},
		{12, []int{1, 2}},
		{123456789, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, c := range testCases {
		//when
		got := getDigits(c.num)
		//then
		if !commonintutils.IntSliceEquals(got, c.want) {
			t.Errorf("getDigits(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

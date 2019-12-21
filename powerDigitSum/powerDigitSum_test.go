package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestPowDigitSum(t *testing.T) {
	testCases := []struct {
		num  float64
		want int
	}{
		//where
		{0, 1},
		{15, 26}, //provided example
	}
	for _, c := range testCases {
		//when
		got := powerDigitSum(c.num)
		//then
		if got != c.want {
			t.Errorf("powDigitSum(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

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

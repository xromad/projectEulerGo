package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestGetProperDivisors(t *testing.T) {
	testCases := []struct {
		num  int64
		want []int64
	}{
		//where
		{1, []int64{}},
		{2, []int64{1}},
		{4, []int64{1, 2}},
		{220, []int64{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110}}, // provided example
		{284, []int64{1, 2, 4, 71, 142}},                        // provided example
	}
	for _, c := range testCases {
		//when
		got := getProperDivisors(c.num)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getProperDivisors(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

func TestGetAmicableNumbers(t *testing.T) {
	testCases := []struct {
		num  int64
		want []int64
	}{
		//where
		{1, []int64{}},
		{30, []int64{}},
		{219, []int64{}},
		{220, []int64{}},
		{283, []int64{}},
		{284, []int64{220, 284}}, // provided example
		{285, []int64{220, 284}}, // provided example
		{286, []int64{220, 284}}, // provided example
	}
	for _, c := range testCases {
		//when
		got := getAmicableNumbers(c.num)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getAmicableNumbers(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

func TestGetAmicableNumbersSum(t *testing.T) {
	testCases := []struct {
		num  int64
		want int64
	}{
		//where
		{1, 0},
		{30, 0},
		{219, 0},
		{220, 0},
		{283, 0},
		{284, 504}, // provided example
		{285, 504}, // provided example
		{286, 504}, // provided example
	}
	for _, c := range testCases {
		//when
		got := getAmicableNumbersSum(c.num)
		//then
		if got != c.want {
			t.Errorf("getAmicableNumbersSum(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestGetFactors(t *testing.T) {
	//given
	testCases := []struct {
		number int64
		want   []int64
	}{
		//where
		{1, []int64{1}},               //perfect
		{2, []int64{1}},               //deficient
		{4, []int64{1, 2}},            //deficient
		{12, []int64{1, 2, 3, 4, 6}},  //provided abundant example
		{28, []int64{1, 2, 4, 7, 14}}, //provided perfect example
	}
	for _, c := range testCases {
		//when
		got := getFactors(c.number)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getFactors(%v) == %v want %v", c.number, got, c.want)
		}
	}
}

func TestGetNumberType(t *testing.T) {
	//given
	testCases := []struct {
		number int64
		want   numType
	}{
		//where
		{1, perfect},   //perfect
		{2, deficient}, //deficient
		{4, deficient}, //deficient
		{12, abundant}, //provided abundant example
		{28, perfect},  //provided perfect example
	}
	for _, c := range testCases {
		//when
		got := getNumberType(c.number)
		//then
		if got != c.want {
			t.Errorf("getNumberType(%v) == %v want %v", c.number, got, c.want)
		}
	}
}

func TestIsSumOfAbundants(t *testing.T) {
	//given
	testCases := []struct {
		abundantNumbers []int64
		number          int64
		want            bool
	}{
		//where
		{[]int64{}, 0, false},
		{[]int64{}, 1, false},
		{[]int64{}, 11, false},
		{[]int64{12}, 24, true},
		{[]int64{12}, 25, false}, //provided example
	}
	for _, c := range testCases {
		//when
		got := isSumOfabundants(c.abundantNumbers, c.number)
		//then
		if got != c.want {
			t.Errorf("isSumOfAbundants(%v) == %v want %v", c.number, got, c.want)
		}
	}
}

func TestGetNonAbundantSums(t *testing.T) {
	//given
	testCases := []struct {
		number int64
		want   int64
	}{
		//where
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
		{23, 276},
		{24, 276},
		{25, 301},
	}
	for _, c := range testCases {
		//when
		got := getNonAbundentSums(c.number)
		//then
		if got != c.want {
			t.Errorf("getNonAbundentSums(%v) == %v want %v", c.number, got, c.want)
		}
	}
}

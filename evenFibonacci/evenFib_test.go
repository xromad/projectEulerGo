package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestGetSum(t *testing.T) {
	testCases := []struct {
		fibSlice []int64
		want     int64
	}{
		{[]int64{}, 0},
		{[]int64{1}, 1},
		{[]int64{1, 2, 3}, 6},
	}

	for _, c := range testCases {
		got := commonintutils.SumInt64Slice(c.fibSlice)
		if got != c.want {
			t.Errorf("genSequence(%v) == %v, want %v", c.fibSlice, got, c.want)
		}
	}
}

func TestGetEvens(t *testing.T) {
	testCases := []struct {
		fibSlice []int64
		want     []int64
	}{
		{[]int64{}, []int64{}},
		{[]int64{1}, []int64{}},
		{[]int64{1, 2, 3}, []int64{2}},
		{[]int64{1, 2, 3, 5, 8}, []int64{2, 8}},
	}

	for _, c := range testCases {
		got := getEvens(c.fibSlice)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("genSequence(%v) == %v, want %v", c.fibSlice, got, c.want)
		}
	}
}

func TestGenSequence(t *testing.T) {
	testCases := []struct {
		maxValue int64
		want     []int64
	}{
		{0, []int64{}},
		{1, []int64{1}},
		{2, []int64{1, 2}},
		{3, []int64{1, 2, 3}},
		{10, []int64{1, 2, 3, 5, 8}},
	}

	for _, c := range testCases {
		got := genSequence(c.maxValue)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("genSequence(%v) == %v, want %v", c.maxValue, got, c.want)
		}
	}
}

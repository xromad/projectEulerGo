package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestGetStringLiteral(t *testing.T) {
	testCases := []struct {
		want int
	}{
		{1000},
	}

	for _, c := range testCases {
		got := getStringLiteral()
		if len(got) != c.want {
			t.Errorf("len(getStringLiteral()) == %v want %v", len(got), c.want)
		}
	}
}

func TestStringToInt64Slice(t *testing.T) {
	testCases := []struct {
		s    string
		want []int64
	}{
		{"", []int64{}},
		{"0", []int64{0}},
		{"1", []int64{1}},
		{"12", []int64{1, 2}},
		{"1234", []int64{1, 2, 3, 4}},
	}

	for _, c := range testCases {
		got := stringToInt64Slice(c.s)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("stringToInt64Slice(%v) == %v want %v", c.s, got, c.want)
		}
	}
}

func TestGetNextSlice(t *testing.T) {
	testCases := []struct {
		start      int
		length     int
		inputSlice []int64
		want       []int64
	}{
		{0, 0, []int64{0, 1, 2, 3, 4}, []int64{}},
		{0, 2, []int64{0, 1, 2, 3, 4}, []int64{0, 1}},
		{2, 0, []int64{0, 1, 2, 3, 4}, []int64{}},
		{0, 5, []int64{0, 1, 2, 3, 4}, []int64{0, 1, 2, 3, 4}},
		{2, 1, []int64{0, 1, 2, 3, 4}, []int64{2}},
		{1, 2, []int64{0, 1, 2, 3, 4}, []int64{1, 2}},
		{2, 3, []int64{0, 1, 2, 3, 4}, []int64{2, 3, 4}},
		{1, 2, []int64{0, 1, 2, 3, 4}, []int64{1, 2}},
		{3, 6, []int64{0, 1, 2, 3, 4}, []int64{3, 4}},
	}

	for _, c := range testCases {
		got := getNextSlice(c.start, c.length, c.inputSlice)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getNextSlice(%v, %v) == %v want %v", c.start, c.length, got, c.want)
		}
	}
}

func TestFindLargestProduct(t *testing.T) {
	testCases := []struct {
		length      int
		inputSlice  []int64
		wantSlice   []int64
		wantProduct int64
	}{
		{2, []int64{0, 1, 2, 3, 4}, []int64{3, 4}, 12},
		{3, []int64{0, 1, 2, 3, 4}, []int64{2, 3, 4}, 24},
	}

	for _, c := range testCases {
		gotSlice, gotProduct := findLargestProduct(c.length, c.inputSlice)
		if !commonintutils.Int64SliceEquals(gotSlice, c.wantSlice) || gotProduct != c.wantProduct {
			t.Errorf("getNextSlice(%v, %v) == %v, %v want %v, %v",
				c.length,
				c.inputSlice,
				gotSlice,
				c.wantSlice,
				gotProduct,
				c.wantProduct,
			)
		}
	}
}

package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestEvenFunction(t *testing.T) {
	testCases := []struct {
		num  int64
		want int64
	}{
		//where
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 1},
	}
	for _, c := range testCases {
		//when
		got := evenFunction(c.num)
		//then
		if got != c.want {
			t.Errorf("evenFunction(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

func TestOddFunction(t *testing.T) {
	testCases := []struct {
		num  int64
		want int64
	}{
		//where
		{0, 1},
		{1, 4},
		{2, 7},
		{3, 10},
	}
	for _, c := range testCases {
		//when
		got := oddFunction(c.num)
		//then
		if got != c.want {
			t.Errorf("oddFunction(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

func TestGenerateSequence(t *testing.T) {
	testCases := []struct {
		num  int64
		want []int64
	}{
		//where
		{1, []int64{1}},
		{2, []int64{2, 1}},
		{13, []int64{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}
	for _, c := range testCases {
		//when
		got := generateSequence(c.num)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("generateSequence(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

func TestFindLongest(t *testing.T) {
	testCases := []struct {
		num  int64
		want []int64
	}{
		//where
		{1, []int64{1}},
		{2, []int64{2, 1}},
		{13, []int64{9, 28, 14, 7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}
	for _, c := range testCases {
		//when
		got := findLongest(c.num)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("findLongest(%v) == %v want %v", c.num, got, c.want)
		}
	}
}

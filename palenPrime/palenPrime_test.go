package main

import (
	"projectEuler/common/commonintutils"
	"testing"
  "reflect"
)

func TestFindEmirps(t *testing.T) {
	testCases := []struct {
		primes []int64
		want   map[int64]int64
	}{
		{[]int64{}, map[int64]int64{}},
		{[]int64{13}, map[int64]int64{}},
		{[]int64{13, 31}, map[int64]int64{13: 31}}, //the provided example
		{[]int64{13, 17}, map[int64]int64{}},
	}

	for _, c := range testCases {
		got := findEmirps(c.primes)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("findEmirps(%v) == %v want %v", c.primes, got, c.want)
		}
	}
}

func TestCheckEmirp(t *testing.T) {
	testCases := []struct {
		primes    []int64
		p         int64
		wantLower int64
		wantUpper int64
		isEmirp   bool
	}{
		{[]int64{}, 0, 0, 0, false},
		{[]int64{}, 13, 0, 0, false},
		{[]int64{13, 31}, 13, 13, 31, true}, //the provided example
		{[]int64{13, 31}, 31, 13, 31, true}, //the provided example
		{[]int64{13, 17}, 13, 0, 0, false},
		{[]int64{131}, 131, 0, 0, false},
	}

	for _, c := range testCases {
		lower, upper, isEmirp := checkEmirp(c.primes, c.p)
		if !((lower == c.wantLower) && (upper == c.wantUpper) && (isEmirp == c.isEmirp)) {
			t.Errorf("checkEmirp(%v, %v) == %v, %v, %t want %v, %v, %t",
				c.primes,
				c.p,
				lower,
				upper,
				isEmirp,
				c.wantLower,
				c.wantUpper,
				c.isEmirp,
			)
		}
	}
}

func TestFindPalendromes(t *testing.T) {
	testCases := []struct {
		nums []int64
		want []int64
	}{
		{[]int64{}, []int64{}},
		{[]int64{101}, []int64{101}}, //the provided example
		{[]int64{121}, []int64{121}},
		{[]int64{123}, []int64{}},
		{[]int64{121, 123}, []int64{121}},
	}

	for _, c := range testCases {
		got := findPalendromes(c.nums)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("findPalendromes(%v) == %v, want %v", c.nums, got, c.want)
		}
	}
}

func TestIsPalendrome(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"aba", true},
		{"abc", false},
		{"101", true}, //the provided example
	}

	for _, c := range testCases {
		got := isPalendrome(c.s)
		if got != c.want {
			t.Errorf("isPalendrome(%s) == %t, want %t", c.s, got, c.want)
		}
	}
}

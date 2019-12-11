package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestGetTriangleNumber(t *testing.T) {
	testCases := []struct {
		n            int
		wantTriangle TriangleNumber
		wantDivisors []int64
	}{
		//where
		{0, TriangleNumber{1, 1}, []int64{1}},
		{1, TriangleNumber{3, 2}, []int64{1, 3}},
		{3, TriangleNumber{6, 3}, []int64{1, 2, 3, 6}},
		{5, TriangleNumber{28, 7}, []int64{1, 2, 4, 7, 14, 28}}, //provided example
	}
	for _, c := range testCases {
		//when
		gotDivisors, gotTriangle := getTriangleNumber(c.n)
		//then
		if !commonintutils.Int64SliceEquals(gotDivisors, c.wantDivisors) {
			t.Errorf("getTriangleNumber(%v) == %v want %v", c.n, gotDivisors, c.wantDivisors)
		}
		if gotTriangle != c.wantTriangle {
			t.Errorf("getTriangleNumber(%v) == %v want %v", c.n, gotTriangle, c.wantTriangle)
		}
	}
}

func TestGetFactors(t *testing.T) {
	testCases := []struct {
		n    int64
		want []int64
	}{
		//where
		{0, []int64{}},
		{1, []int64{1}},
		{3, []int64{1, 3}},
		{6, []int64{1, 2, 3, 6}},
		{10, []int64{1, 2, 5, 10}},
		{15, []int64{1, 3, 5, 15}},
		{21, []int64{1, 3, 7, 21}},
		{28, []int64{1, 2, 4, 7, 14, 28}}, //provided example
	}
	for _, c := range testCases {
		//when
		got := getFactors(c.n)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getFactors(%v) == %v want %v", c.n, got, c.want)
		}
	}
}

func TestGetNextTriangle(t *testing.T) {
	testCases := []struct {
		triangleNumber TriangleNumber
		want           TriangleNumber
	}{
		//where
		{TriangleNumber{0, 0}, TriangleNumber{1, 1}},
		{TriangleNumber{1, 1}, TriangleNumber{3, 2}},
		{TriangleNumber{3, 2}, TriangleNumber{6, 3}},
	}
	for _, c := range testCases {
		//when
		got := getNextTriangle(c.triangleNumber)
		//then
		if !triangleNumberEquals(got, c.want) {
			t.Errorf("getNextTriangle(%v) == %v want %v", c.triangleNumber, got, c.want)
		}
	}
}

//Test helper Function
func triangleNumberEquals(t1 TriangleNumber, t2 TriangleNumber) bool {
	if t1.value == t2.value && t1.lastInt == t2.lastInt {
		return true
	}
	return false
}

//Test of test helper function
func TestTriangleNumberEquals(t *testing.T) {
	testCases := []struct {
		t1   TriangleNumber
		t2   TriangleNumber
		want bool
	}{
		//where
		{TriangleNumber{0, 0}, TriangleNumber{0, 0}, true},
		{TriangleNumber{1, 0}, TriangleNumber{0, 0}, false},
		{TriangleNumber{0, 1}, TriangleNumber{0, 0}, false},
		{TriangleNumber{0, 0}, TriangleNumber{1, 0}, false},
		{TriangleNumber{0, 0}, TriangleNumber{0, 1}, false},
		{TriangleNumber{1, 1}, TriangleNumber{0, 0}, false},
		{TriangleNumber{0, 0}, TriangleNumber{1, 1}, false},
		{TriangleNumber{1, 1}, TriangleNumber{1, 1}, true},
	}
	for _, c := range testCases {
		//when
		got := triangleNumberEquals(c.t1, c.t2)
		//then
		if got != c.want {
			t.Errorf("triangleNumberEqals(%v, %v) == %v want %v", c.t1, c.t2, got, c.want)
		}
	}
}

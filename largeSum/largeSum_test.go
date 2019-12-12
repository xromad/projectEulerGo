package main

import (
	"math/big"
	"testing"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		a    int
		b    int
		want int
	}{
		//where
		{0, 0, 0},
		{1, 1, 1},
		{1, 2, 1},
		{2, 1, 1},
	}
	for _, c := range testCases {
		//when
		got := min(c.a, c.b)
		//then
		if got != c.want {
			t.Errorf("min(%v, %v) == %v want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestSumNumArray(t *testing.T) {
	testCases := []struct {
		numArray []string
		want     *big.Int
	}{
		//where
		{[]string{"1", "2", "3", "4", "5"}, big.NewInt(15)},
		{[]string{"0", "0", "0", "0", "0"}, big.NewInt(0)},
		{[]string{"15"}, big.NewInt(15)},
	}
	for _, c := range testCases {
		//when
		got := sumNumArray(c.numArray)
		//then
		if got.Cmp(c.want) != 0 {
			t.Errorf("sumNumArray(%v) == %v want %v", c.numArray, got, c.want)
		}
	}
}

func TestGenerateNumArray(t *testing.T) {
	testCases := []struct {
		wantFirst string
		wantLast  string
		wantSize  int
	}{
		//where
		{
			"37107287533902102798797998220837590246510135740250",
			"53503534226472524250874054075591789781264330331690",
			100,
		},
	}
	for _, c := range testCases {
		//when
		got := generateNumArray()
		//then
		if len(got) != c.wantSize {
			t.Errorf("size(generateNumArray()) == %v want %v", len(got), c.wantSize)
		}
		//and
		if got[0] != c.wantFirst {
			t.Errorf("first(generateNumArray()) == %v want %v", got[0], c.wantFirst)
		}
		//and
		if got[len(got)-1] != c.wantLast {
			t.Errorf("generateNumArray() == %v want %v", got[len(got)-1], c.wantLast)
		}
	}
}

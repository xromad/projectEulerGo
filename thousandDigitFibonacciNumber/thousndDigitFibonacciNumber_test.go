package main

import (
	"math/big"
	"testing"
)

func TestGetIndex(t *testing.T) {
	//given
	testCases := []struct {
		count int
		want  fibonacciNumber
	}{
		//where
		{2, fibonacciNumber{*big.NewInt(int64(13)), 7}},
		{3, fibonacciNumber{*big.NewInt(int64(144)), 12}},
	}
	for _, c := range testCases {
		//when
		got := getIndex(c.count)
		//then
		if got.num.String() != c.want.num.String() || got.index != c.want.index {
			t.Errorf("getIndex(%v) == f[%v,%v] want f[%v,%v]",
				c.count,
				got.num.String(), got.index,
				c.want.num.String(), c.want.index)
		}
	}
}

func TestGetNthLexiPermutation(t *testing.T) {
	//given
	testCases := []struct {
		previousFibNumIn fibonacciNumber
		fibNumIn         fibonacciNumber
		want             fibonacciNumber
	}{
		//where
		{fibonacciNumber{*big.NewInt(int64(1)), 1}, fibonacciNumber{*big.NewInt(int64(1)), 2}, fibonacciNumber{*big.NewInt(int64(2)), 3}},        // F3 = 2
		{fibonacciNumber{*big.NewInt(int64(1)), 2}, fibonacciNumber{*big.NewInt(int64(2)), 3}, fibonacciNumber{*big.NewInt(int64(3)), 4}},        // F4 = 3
		{fibonacciNumber{*big.NewInt(int64(2)), 3}, fibonacciNumber{*big.NewInt(int64(3)), 4}, fibonacciNumber{*big.NewInt(int64(5)), 5}},        // F5 = 5
		{fibonacciNumber{*big.NewInt(int64(3)), 4}, fibonacciNumber{*big.NewInt(int64(5)), 5}, fibonacciNumber{*big.NewInt(int64(8)), 6}},        // F6 = 8
		{fibonacciNumber{*big.NewInt(int64(5)), 5}, fibonacciNumber{*big.NewInt(int64(8)), 6}, fibonacciNumber{*big.NewInt(int64(13)), 7}},       // F7 = 13
		{fibonacciNumber{*big.NewInt(int64(8)), 6}, fibonacciNumber{*big.NewInt(int64(13)), 7}, fibonacciNumber{*big.NewInt(int64(21)), 8}},      // F8 = 21
		{fibonacciNumber{*big.NewInt(int64(13)), 7}, fibonacciNumber{*big.NewInt(int64(21)), 8}, fibonacciNumber{*big.NewInt(int64(34)), 9}},     // F9 = 34
		{fibonacciNumber{*big.NewInt(int64(21)), 8}, fibonacciNumber{*big.NewInt(int64(34)), 9}, fibonacciNumber{*big.NewInt(int64(55)), 10}},    // F10 = 55
		{fibonacciNumber{*big.NewInt(int64(34)), 9}, fibonacciNumber{*big.NewInt(int64(55)), 10}, fibonacciNumber{*big.NewInt(int64(89)), 11}},   // F11 = 89
		{fibonacciNumber{*big.NewInt(int64(55)), 10}, fibonacciNumber{*big.NewInt(int64(89)), 11}, fibonacciNumber{*big.NewInt(int64(144)), 12}}, // F12 = 144
	}
	for _, c := range testCases {
		//when
		got, _ := getNextFibonacciNumber(c.previousFibNumIn, c.fibNumIn)
		//then
		if got.num.String() != c.want.num.String() || got.index != c.want.index {
			t.Errorf("getNextFibonacciNumber(f[%v,%v]) == f[%v,%v] want f[%v,%v]",
				c.fibNumIn.num.String(), c.fibNumIn.index,
				got.num.String(), got.index,
				c.want.num.String(), c.want.index)
		}
	}
}

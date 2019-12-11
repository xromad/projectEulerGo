package main

import (
	"testing"
)

func TestGenOutput(t *testing.T) {
	testCases := []struct {
		start int
		end   int
		want  string
	}{
		{10, 15, "buzz 11 fizz 13 14 fizzbuzz"},
		{1, 2, "1 2"},
		{0, 0, "fizzbuzz"},
		{0, 1, "fizzbuzz 1"},
	}

	for _, c := range testCases {
		got := genOutput(c.start, c.end)
		if got != c.want {
			t.Errorf("genOutput(%d, %d) == %s, want %s", c.start, c.end, got, c.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{15, "fizzbuzz"},
		{5, "buzz"},
		{3, "fizz"},
		{4, "4"},
		{0, "fizzbuzz"},
	}

	for _, c := range testCases {
		got := getFizzBuzz(c.num)
		if got != c.want {
			t.Errorf("getFizzBuzz(%d) == %s, want %s", c.num, got, c.want)
		}
	}
}

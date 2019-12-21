package main

import (
	"testing"
)

func TestGetNumberLetterCounts(t *testing.T) {
	testCases := []struct {
		max       int
		wantWords string
		wantCount int
	}{
		//where
		{1, "one", 3},
		{2, "one two", 6},
		{5, "one two three four five", 19},
	}
	for _, c := range testCases {
		//when
		gotWords, gotCount := getNumberLetterCounts(c.max)
		//then
		if gotWords != c.wantWords {
			t.Errorf("processNums(%v) == %s want %s", c.max, gotWords, c.wantWords)
		}
		//and
		if gotCount != c.wantCount {
			t.Errorf("processNums(%v) == %v want %v", c.max, gotCount, c.wantCount)
		}
	}
}

func TestProcessNums(t *testing.T) {
	testCases := []struct {
		max       int
		wantWords string
		wantCount int
	}{
		//where
		{1, "one", 3},
		{2, "one two", 6},
		{5, "one two three four five", 19},
	}
	for _, c := range testCases {
		//when
		gotWords, gotCount := processNums(c.max)
		//then
		if gotWords != c.wantWords {
			t.Errorf("processNums(%v) == %s want %s", c.max, gotWords, c.wantWords)
		}
		//and
		if gotCount != c.wantCount {
			t.Errorf("processNums(%v) == %v want %v", c.max, gotCount, c.wantCount)
		}
	}
}

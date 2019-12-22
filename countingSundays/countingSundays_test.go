package main

import (
	"testing"
)

func TestCountSundays(t *testing.T) {
	testCases := []struct {
		startDate string
		endDate   string
		want      int
	}{
		//where
		{"Jan-01-1901", "Dec-31-1901", 51},
	}
	for _, c := range testCases {
		//when
		got := countSundays(c.startDate, c.endDate)
		//then
		if got != c.want {
			t.Errorf("countSundays(%s, %s) == %v want %v", c.startDate, c.endDate, got, c.want)
		}
	}
}

func TestDaysToSunday(t *testing.T) {
	testCases := []struct {
		day  string
		want int
	}{
		//where
		{"Sunday", 0},
		{"Monday", 6},
		{"Tuesday", 5},
		{"Wednesday", 4},
		{"Thursday", 3},
		{"Friday", 2},
		{"Saturday", 1},
	}
	for _, c := range testCases {
		//when
		got := daysToSunday(c.day)
		//then
		if got != c.want {
			t.Errorf("daysToSunday(%s) == %v want %v", c.day, got, c.want)
		}
	}
}

package main

import (
	"testing"
)

func TestGetNamesScore(t *testing.T) {
	//given
	testCases := []struct {
		fileName string
		want     int64
	}{
		//where
		{"p022_names.txt", 871198282},
	}
	for _, c := range testCases {
		//when
		got := getNamesScore(c.fileName)
		//then
		if got != c.want {
			t.Errorf("getNamesScore(%s) == %v want %v", c.fileName, got, c.want)
		}
	}
}

func TestGetTotalScore(t *testing.T) {
	//given
	testCases := []struct {
		names []string
		want  int64
	}{
		//where
		{[]string{"\"A\"", "\"A\"", "\"A\""}, 6},
		{[]string{"\"AA\"", "\"AA\"", "\"AA\""}, 12},
		{[]string{"\"B\"", "\"B\"", "\"B\""}, 12},
		{[]string{"\"BB\"", "\"BB\"", "\"BB\""}, 24},
		{[]string{"\"A\"", "\"B\"", "\"C\""}, 14},
		{[]string{"\"a\"", "\"b\"", "\"c\""}, 14},
		{[]string{"\"Aa\"", "\"Bb\"", "\"Cc\""}, 28},
		{[]string{"\"Jenny\"", "\"Brent\""}, 186},
	}
	for _, c := range testCases {
		//when
		got := getTotalScore(c.names)
		//then
		if got != c.want {
			t.Errorf("getTotalScore(%v) == %v want %v", c.names, got, c.want)
		}
	}
}

func TestGetScoreForName(t *testing.T) {
	//given
	testCases := []struct {
		name  string
		index int64
		want  int64
	}{
		//where
		{"\"A\"", 1, 1},
		{"\"AA\"", 2, 4},
		{"\"B\"", 1, 2},
		{"\"B\"", 2, 4},
		{"\"B\"", 3, 6},
		{"\"Z\"", 1, 26},
		{"\"z\"", 10, 260},
		{"\"Brent\"", 10, 590},
		{"\"Jenny\"", 10, 680},
	}
	for _, c := range testCases {
		//when
		got := getScoreForName(c.index, c.name)
		//then
		if got != c.want {
			t.Errorf("getScoreForName(%v, %s) == %v want %v", c.index, c.name, got, c.want)
		}
	}
}

func TestGetAlpha(t *testing.T) {
	//given
	testCases := []struct {
		character string
		want      int64
	}{
		//where
		{"A", 1},
		{"a", 1},
		{"Z", 26},
		{"z", 26},
	}
	for _, c := range testCases {
		//when
		got := getAlpha(c.character)
		//then
		if got != c.want {
			t.Errorf("getAlpha(%s) == %v want %v", c.character, got, c.want)
		}
	}
}

func TestParseFile(t *testing.T) {
	//given
	testCases := []struct {
		fileName  string
		arraySize int
		firstName string
		lastName  string
	}{
		//where
		{"p022_names.txt", 5163, "\"MARY\"", "\"ALONSO\""}, // the provided example
	}
	for _, c := range testCases {
		//when
		got := parseFile(c.fileName)
		//then
		if len(got) != c.arraySize {
			t.Errorf("parseFile(%v): arraySize = %v want %v", c.fileName, len(got), c.arraySize)
		}
		//and
		if got[0] != c.firstName {
			t.Errorf("parseFile(%v): firstName == %v want %v", c.fileName, got[0], c.firstName)
		}
		//and
		if got[len(got)-1] != c.lastName {
			t.Errorf("parseFile(%v): firstName == %v want %v", c.fileName, got[len(got)-1], c.lastName)
		}
	}
}

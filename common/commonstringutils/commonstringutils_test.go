package commonstringutils

import (
	"testing"
)

func TestIsPalendrome(t *testing.T) {
	testCases := []struct {
		sIn  string
		want bool
	}{
		{"", true}, //empty
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"£¤£", true}, //unicode
	}

	for _, c := range testCases {
		got := IsPalendrome(c.sIn)
		if got != c.want {
			t.Errorf("reverseString(%s) == %t, want %t", c.sIn, got, c.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		sIn  string
		want string
	}{
		{"", ""}, //empty
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"£¤¥", "¥¤£"}, //unicode
	}

	for _, c := range testCases {
		got := ReverseString(c.sIn)
		if got != c.want {
			t.Errorf("reverseString(%s) == %s, want %s", c.sIn, got, c.want)
		}
	}
}

func TestStringSliceCompare(t *testing.T) {
	testCases := []struct {
		sliceA []string
		sliceB []string
		want   bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"a"}, []string{}, false},
		{[]string{}, []string{"b"}, false},
		{[]string{"a"}, []string{"b"}, false},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, false},
		{[]string{"a"}, []string{"a", "b"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
	}

	for _, c := range testCases {
		got := StringSliceCompare(c.sliceA, c.sliceB)
		if got != c.want {
			t.Errorf("StringSliceCompare(%v, %v) == %v, want %v", c.sliceA, c.sliceB, got, c.want)
		}
	}
}

func TestStringHasVals(t *testing.T) {
	testCases := []struct {
		sliceA []string
		sliceB []string
		want   bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"a"}, []string{}, false},
		{[]string{}, []string{"b"}, false},
		{[]string{"a"}, []string{"b"}, false},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, false},
		{[]string{"a"}, []string{"a", "b"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
	}

	for _, c := range testCases {
		got := StringHasVals(c.sliceA, c.sliceB)
		if got != c.want {
			t.Errorf("StringHasVals(%v, %v) == %v, want %v", c.sliceA, c.sliceB, got, c.want)
		}
	}
}

func TestStringSliceContains(t *testing.T) {
	testCases := []struct {
		sliceA []string
		s      string
		want   bool
	}{
		{[]string{}, "a", false},
		{[]string{"a"}, "a", true},
		{[]string{"a"}, "b", false},
		{[]string{"a", "b"}, "a", true},
		{[]string{"a", "b"}, "b", true},
	}

	for _, c := range testCases {
		got := StringSliceContains(c.sliceA, c.s)
		if got != c.want {
			t.Errorf("StringSliceContains(%v, %v) == %v, want %v", c.sliceA, c.s, got, c.want)
		}
	}
}

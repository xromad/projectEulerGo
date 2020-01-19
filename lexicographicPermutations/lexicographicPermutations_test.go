package main

import (
	"projectEuler/common/commonstringutils"
	"testing"
)

func TestGetNthLexiPermutation(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn      []string
		count        int64
		wantMutation []string
		wantCount    int64
	}{
		//where
		{[]string{}, 1, []string{}, 1},
		{[]string{"1"}, 1, []string{"1"}, 1},
		{[]string{"1", "2"}, 1, []string{"1", "2"}, 1},
		{[]string{"1", "2"}, 2, []string{"2", "1"}, 2},
		{[]string{"0", "1", "2"}, 1, []string{"0", "1", "2"}, 1}, //provided example
		{[]string{"0", "1", "2"}, 2, []string{"0", "2", "1"}, 2}, //provided example
		{[]string{"0", "1", "2"}, 3, []string{"1", "0", "2"}, 3}, //provided example
		{[]string{"0", "1", "2"}, 4, []string{"1", "2", "0"}, 4}, //provided example
		{[]string{"0", "1", "2"}, 5, []string{"2", "0", "1"}, 5}, //provided example
		{[]string{"0", "1", "2"}, 6, []string{"2", "1", "0"}, 6}, //provided example
		{[]string{"0", "1", "2"}, 7, []string{"2", "1", "0"}, 6}, //provided example
	}
	for _, c := range testCases {
		//when
		gotMutation, gotCount := getNthLexiPermutation(c.sliceIn, c.count)
		//then
		if !commonstringutils.StringSliceCompare(gotMutation, c.wantMutation) || gotCount != c.wantCount {
			t.Errorf("getNthLexiPermutation(%s, %v) == %s,%v want %s, %v",
				c.sliceIn, c.count, gotMutation, gotCount, c.wantMutation, c.wantCount)
		}
	}
}

func TestStringSliceExactMatch(t *testing.T) {
	//given
	testCases := []struct {
		sliceA []string
		sliceB []string
		want   bool
	}{
		//where
		{[]string{}, []string{}, true},
		{[]string{"1"}, []string{"1"}, true},
		{[]string{"1", "2"}, []string{"1", "1"}, false},
		{[]string{"1", "2"}, []string{"2", "2"}, false},
		{[]string{"1", "2"}, []string{"2", "1"}, false},
		{[]string{"1", "1"}, []string{"1", "1"}, true},
		{[]string{"2", "2"}, []string{"2", "2"}, true},
	}
	for _, c := range testCases {
		//when
		got := stringSliceExactMatch(c.sliceA, c.sliceB)
		//then
		if got != c.want {
			t.Errorf("getNextMutation(%s, %s) == %v, want %v", c.sliceA, c.sliceB, got, c.want)
		}
	}
}

func TestGetNextMutation(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn []string
		want    []string
	}{
		//where
		{[]string{}, []string{}},
		{[]string{"1"}, []string{"1"}},
		{[]string{"1", "2"}, []string{"2", "1"}},
		{[]string{"0", "1", "2"}, []string{"0", "2", "1"}}, //provided example
		{[]string{"0", "2", "1"}, []string{"1", "0", "2"}}, //provided example
		{[]string{"1", "0", "2"}, []string{"1", "2", "0"}}, //provided example
		{[]string{"1", "2", "0"}, []string{"2", "0", "1"}}, //provided example
		{[]string{"2", "0", "1"}, []string{"2", "1", "0"}}, //provided example
	}
	for _, c := range testCases {
		//when
		got := getNextMutation(c.sliceIn)
		//then
		if !commonstringutils.StringSliceCompare(got, c.want) {
			t.Errorf("getNextMutation(%s) == %s, want %s", c.sliceIn, got, c.want)
		}
	}
}

func TestFindNextSmallestRightCharacter(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn       []string
		indexF        int
		wantCharacter string
		wantIndex     int
	}{
		//where
		{[]string{}, 1, "", 0},
		{[]string{"1"}, 0, "1", 0},
		{[]string{"1", "2"}, 0, "2", 1},
		{[]string{"1", "2", "3", "4", "5"}, 3, "5", 4},
		{[]string{"1", "2", "3", "5", "4"}, 2, "4", 4},
		{[]string{"1", "2", "3", "4", "5"}, 2, "4", 3},
	}
	for _, c := range testCases {
		//when
		gotCharacter, gotIndex := findNextSmallestRightCharacter(c.sliceIn, c.indexF)
		//then
		if gotCharacter != c.wantCharacter || gotIndex != c.wantIndex {
			t.Errorf("findNextSmallestRightCharacter(%s) == %s, %v want %s, %v", c.sliceIn, gotCharacter, gotIndex, c.wantCharacter, c.wantIndex)
		}
	}
}

func TestFindLastSmallestCharacter(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn       []string
		wantCharacter string
		wantIndex     int
	}{
		//where
		{[]string{}, "", 0},
		{[]string{"1"}, "1", 0},
		{[]string{"1", "2"}, "1", 0},
		{[]string{"1", "2", "3", "4", "5"}, "4", 3},
		{[]string{"1", "2", "3", "5", "4"}, "3", 2},
	}
	for _, c := range testCases {
		//when
		gotCharacter, gotIndex := findLastSmallestCharacter(c.sliceIn)
		//then
		if gotCharacter != c.wantCharacter || gotIndex != c.wantIndex {
			t.Errorf("findLastSmallestCharacter(%s) == %s, %v want %s, %v", c.sliceIn, gotCharacter, gotIndex, c.wantCharacter, c.wantIndex)
		}
	}
}

func TestSortAfter(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn []string
		index   int
		want    []string
	}{
		//where
		{[]string{}, 2, []string{}},
		{[]string{"1"}, 0, []string{"1"}},
		{[]string{"1", "2"}, 0, []string{"1", "2"}},
		{[]string{"1", "2", "4", "5", "3"}, 2, []string{"1", "2", "4", "3", "5"}},
	}
	for _, c := range testCases {
		//when
		got := sortAfter(c.sliceIn, c.index)
		//then
		if !commonstringutils.StringSliceCompare(got, c.want) {
			t.Errorf("sortAfter(%s, %v) == %v want %v", c.sliceIn, c.index, got, c.want)
		}
	}
}

func TestSwapCharacters(t *testing.T) {
	//given
	testCases := []struct {
		sliceIn []string
		indexF  int
		indexS  int
		want    []string
	}{
		//where
		{[]string{}, 2, 4, []string{}},
		{[]string{"1"}, 0, 0, []string{"1"}},
		{[]string{"1", "2"}, 0, 1, []string{"2", "1"}},
		{[]string{"1", "2", "3", "4", "5"}, 1, 3, []string{"1", "4", "3", "2", "5"}},
	}
	for _, c := range testCases {
		//when
		got := swapCharacters(c.sliceIn, c.indexF, c.indexS)
		//then
		if !commonstringutils.StringSliceCompare(got, c.want) {
			t.Errorf("swapCharacters(%s, %v, %v) == %v want %v", c.sliceIn, c.indexF, c.indexS, got, c.want)
		}
	}
}

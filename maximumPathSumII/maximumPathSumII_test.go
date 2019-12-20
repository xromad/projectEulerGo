package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestMaxPath(t *testing.T) {
	//given
	tree, _ := makeTree(true)

	testCases := []struct {
		want uint
	}{
		//where
		{23}, // the provided example
	}
	for _, c := range testCases {
		//when
		got := checkTree(tree)
		//then
		if got != c.want {
			t.Errorf("checkTree(%v) == %v want %v", tree, got, c.want)
		}
	}
}

func TestGetNodes(t *testing.T) {
	//given
	tree, _ := makeTree(true)

	testCases := []struct {
		want []uint
	}{
		//where
		{[]uint{3, 7, 4, 9}},
	}
	for _, c := range testCases {
		//when
		got := getNodes(tree)
		//then
		if commonintutils.UIntSliceEquals(got, c.want) {
			t.Errorf("sumNodes(%v) == %v want %v", tree, got, c.want)
		}
	}
}

func TestMakeTree(t *testing.T) {
	testCases := []struct {
		sample   bool
		wantVal  uint
		wantRows uint
	}{
		//where
		{true, 3, 4},
		{false, 59, 100},
	}

	for _, c := range testCases {
		//when
		gotTree, gotRows := makeTree(c.sample)
		//then
		if gotTree.value != c.wantVal {
			t.Errorf("makeTree(%v) == %v value want %v ", c.sample, gotTree.value, c.wantVal)
		}
		//and
		if gotRows != c.wantRows {
			t.Errorf("makeTree(%v) == %v rows want %v ", c.sample, gotRows, c.wantRows)
		}
	}
}

func TestParseTree(t *testing.T) {
	testCases := []struct {
		filenName string
		wantLeft  uint
		wantRight uint
		wantRows  uint
	}{
		//where
		{"sampleTriangle.txt", 7, 4, 4},
		{"p067_triangle.txt", 73, 41, 100},
	}

	for _, c := range testCases {
		//when
		tree, rows := parseTree(c.filenName)
		//then
		gotLeft := tree.left.value
		if gotLeft != c.wantLeft {
			t.Errorf("makeExampleTree(%s) == %v, left == %v want %v in ", c.filenName, tree, gotLeft, c.wantLeft)
		}
		//and
		gotRight := tree.right.value
		if gotRight != c.wantRight {
			t.Errorf("makeExampleTree(%s) == %v, right == %v want %v in ", c.filenName, tree, gotRight, c.wantRight)
		}
		//and
		if rows != c.wantRows {
			t.Errorf("makeExampleTree(%s) == %v, rows == %v want %v in ", c.filenName, tree, rows, c.wantRows)
		}
	}
}

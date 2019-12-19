package main

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestMaxPath(t *testing.T) {
	//given
	tree := makeExampleTree()

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
	tree := makeExampleTree()

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
		sample bool
		want   uint
	}{
		//where
		{true, 3},
		{false, 75},
	}

	for _, c := range testCases {
		//when
		got := makeTree(c.sample)
		//then
		if got.value != c.want {
			t.Errorf("makeTree(%v) == %v want %v ", c.sample, got.value, c.want)
		}
	}
}

func TestMakeExampleTree(t *testing.T) {
	testCases := []struct {
		wantLeft  uint
		wantRight uint
	}{
		//where
		{7, 4},
	}

	for _, c := range testCases {
		//when
		tree := makeExampleTree()
		//then
		gotLeft := tree.left.value
		if gotLeft != c.wantLeft {
			t.Errorf("makeExampleTree() == %v, left == %v want %v in ", tree, gotLeft, c.wantLeft)
		}
		//and
		gotRight := tree.right.value
		if gotRight != c.wantRight {
			t.Errorf("makeExampleTree() == %v, right == %v want %v in ", tree, gotRight, c.wantRight)
		}
	}
}

func TestMakeRealTree(t *testing.T) {
	testCases := []struct {
		wantLeft  uint
		wantRight uint
	}{
		//where
		{95, 64},
	}

	for _, c := range testCases {
		//when
		tree := makeRealTree()
		//then
		gotLeft := tree.left.value
		if gotLeft != c.wantLeft {
			t.Errorf("makeExampleTree() == %v, left == %v want %v in ", tree, gotLeft, c.wantLeft)
		}
		//and
		gotRight := tree.right.value
		if gotRight != c.wantRight {
			t.Errorf("makeExampleTree() == %v, right == %v want %v in ", tree, gotRight, c.wantRight)
		}
	}
}

package main

import (
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
		maxPath := checkTree(tree)
		got := sumNodes(maxPath)
		//then
		if got != c.want {
			t.Errorf("checkTree(%v) == %v want %v", tree, got, c.want)
		}
	}
}

func TestSumNodes(t *testing.T) {
	testCases := []struct {
		path []*node
		want uint
	}{
		//where
		{[]*node{&node{0, nil, nil}}, 0},
		{[]*node{&node{1, nil, nil}}, 1},
		{[]*node{&node{1, nil, nil}, &node{2, nil, nil}, &node{3, nil, nil}}, 6},
	}
	for _, c := range testCases {
		//when
		got := sumNodes(c.path)
		//then
		if got != c.want {
			t.Errorf("sumNodes(%v) == %v want %v", c.path, got, c.want)
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

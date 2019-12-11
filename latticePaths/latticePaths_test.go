package main

import (
	"fmt"
	"testing"
)

func TestIntiGrid(t *testing.T) {
	testCases := []struct {
		gridSize  int
		wantRight int
		wantDown  int
		wantValue uint64
	}{
		{1, 0, 0, 1},
		{2, 1, 1, 1},
	}

	for _, c := range testCases {
		got := initGrid(c.gridSize)
		gotValue := (*got)[c.wantRight][c.wantDown]
		if gotValue != c.wantValue {
			t.Errorf("InitGrid(%v) == %v want %v", c.gridSize, gotValue, c.wantValue)
		}
	}
}

func TestCheckDown(t *testing.T) {
	testCases := []struct {
		gridSize  int
		gridRight int
		gridDown  int
		wantValue uint64
	}{
		{1, 1, 1, 1},
		{2, 1, 2, 1},
	}

	for _, c := range testCases {
		grid := initGrid(c.gridSize)
		checkRight(c.gridRight, c.gridDown, grid)
		fmt.Println(fmt.Sprintf("filled grid: %v", grid))
		gotValue := (*grid)[c.gridRight-1][c.gridDown-1]
		if gotValue != c.wantValue {
			t.Errorf("checkDown(%v) == %v want %v", c.gridSize, gotValue, c.wantValue)
		}
	}
}

func TestFillGrid(t *testing.T) {
	testCases := []struct {
		gridSize  int
		wantValue uint64
	}{
		{1, 1},
		{2, 2},
		{3, 6},
	}

	for _, c := range testCases {
		grid := initGrid(c.gridSize)
		fillGrid(c.gridSize, c.gridSize, c.gridSize, grid)
		fmt.Println(fmt.Sprintf("filled grid: %v", grid))
		gotValue := (*grid)[0][0]
		if gotValue != c.wantValue {
			t.Errorf("fillGrid(%v) == %v want %v", c.gridSize, gotValue, c.wantValue)
		}
	}
}

func TestEasyWay(t *testing.T) {
	testCases := []struct {
		gridSize int
		want     uint64
	}{
		{0, 1},
		{1, 2},
		{2, 6},
	}

	for _, c := range testCases {
		got := easyWay(c.gridSize)
		if got != c.want {
			t.Errorf("easyWay(%v) == %v want %v", c.gridSize, got, c.want)
		}
	}
}

func TestCheckRight(t *testing.T) {
	testCases := []struct {
		gridSize  int
		gridRight int
		gridDown  int
		wantValue uint64
	}{
		{1, 1, 1, 1},
		{2, 2, 1, 1},
	}

	for _, c := range testCases {
		grid := initGrid(c.gridSize)
		checkDown(c.gridRight, c.gridDown, grid)
		fmt.Println(fmt.Sprintf("filled grid: %v", grid))
		gotValue := (*grid)[c.gridRight-1][c.gridDown-1]
		if gotValue != c.wantValue {
			t.Errorf("checkDown(%v) == %v want %v", c.gridSize, gotValue, c.wantValue)
		}
	}
}

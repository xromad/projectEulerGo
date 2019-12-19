package main

import (
	"fmt"
	"projectEuler/common/commonintutils"
	"reflect"
	"testing"
)

func TestSlash(t *testing.T) {

	//given
	givenGrid := [][]int64{
		[]int64{1, 1, 1, 1},
		[]int64{1, 4, 2, 1},
		[]int64{1, 3, 5, 1},
		[]int64{1, 1, 1, 1},
	}

	testCases := []struct {
		direction   string
		count       int
		wantFactors []int64
		wantProduct int64
	}{
		//where
		{"right", 0, []int64{}, 0},
		{"right", 2, []int64{2, 3}, 6},
		{"left", 0, []int64{}, 0},
		{"left", 2, []int64{4, 5}, 20},
	}
	for _, c := range testCases {
		//when
		gotFactors, gotProduct := slash(givenGrid, c.count, c.direction)
		//then
		if !commonintutils.Int64SliceEquals(gotFactors, c.wantFactors) {
			t.Errorf("slashRight(%v, %v, %v) == %v want %v",
				givenGrid,
				c.count,
				c.direction,
				gotFactors,
				c.wantFactors,
			)
		}
		//and
		if gotProduct != c.wantProduct {
			t.Errorf("slashRight(%v, %v) == %v want %v",
				givenGrid,
				c.count,
				gotProduct,
				c.wantProduct,
			)
		}
	}
}

func TestNextRightDiagonal(t *testing.T) {
	// 3X3 grid indicees
	// 	{0:0, 0:1, 0:2}
	// 	{1:0, 1:1, 1:2}
	// 	{2:0, 2:1, 2:2}

	//given
	testCases := []struct {
		diagonal []Cell
		rows     int
		columns  int
		want     []Cell
	}{
		//where
		{[]Cell{}, 3, 3, []Cell{Cell{0, 0}}},
		{[]Cell{Cell{0, 0}}, 3, 3, []Cell{Cell{0, 1}, Cell{1, 0}}},
		{[]Cell{Cell{0, 1}, Cell{1, 0}}, 3, 3, []Cell{Cell{0, 2}, Cell{1, 1}, Cell{2, 0}}},
		{[]Cell{Cell{0, 2}, Cell{1, 1}, Cell{2, 0}}, 3, 3, []Cell{Cell{1, 2}, Cell{2, 1}}},
		{[]Cell{Cell{1, 2}, Cell{2, 1}}, 3, 3, []Cell{Cell{2, 2}}},
	}
	for _, c := range testCases {
		//when
		got := nextRightDiagonal(c.diagonal, c.rows, c.columns)
		//then
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("nextRightDiagonal(%v, %v, %v) == %v want %v",
				c.diagonal,
				c.rows,
				c.columns,
				got,
				c.want,
			)
		}
	}
}

func TestGetMaxKey(t *testing.T) {
	//given
	testCases := []struct {
		diagonal []Cell
		want     int
	}{
		//where
		{[]Cell{}, 0},
		{[]Cell{Cell{0, 0}}, 0},
		{[]Cell{
			Cell{0, 1},
			Cell{1, 0}},
			1},
		{[]Cell{
			Cell{0, 2},
			Cell{1, 1},
			Cell{2, 0}},
			2},
	}
	for _, c := range testCases {
		//when
		got := getMaxRow(c.diagonal)
		//then
		if got != c.want {
			t.Errorf("getMaxKey(%v) == %v want %v",
				c.diagonal,
				got,
				c.want,
			)
		}
	}
}

func TestGetDiagonalSlice(t *testing.T) {

	//given
	givenGrid := [][]int64{
		[]int64{1, 2, 3, 4},
		[]int64{5, 1, 2, 3},
		[]int64{6, 5, 1, 2},
		[]int64{7, 6, 5, 1},
	}

	testCases := []struct {
		diagonal []Cell
		want     []int64
	}{
		//where
		{[]Cell{}, []int64{}},
		{[]Cell{
			Cell{0, 0}},
			[]int64{1}},
		{[]Cell{
			Cell{0, 0},
			Cell{1, 1},
			Cell{2, 2},
			Cell{3, 3}},
			[]int64{1, 1, 1, 1}},
		{[]Cell{
			Cell{0, 3},
			Cell{1, 2},
			Cell{2, 1},
			Cell{3, 0}},
			[]int64{4, 2, 5, 7}},
	}
	for _, c := range testCases {
		//when
		got := getDiagonalSlice(c.diagonal, givenGrid)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("getDiagonalSlice(%v, %v) == %v want %v",
				c.diagonal,
				givenGrid,
				got,
				c.want,
			)
		}
	}
}

func TestRowRight(t *testing.T) {

	//given
	givenGrid := [][]int64{
		[]int64{1, 1, 1, 1},
		[]int64{1, 2, 3, 1},
		[]int64{1, 1, 1, 1},
		[]int64{1, 1, 1, 1},
	}

	testCases := []struct {
		count       int
		wantFactors []int64
		wantProduct int64
	}{
		//where
		{0, []int64{}, 0},
		{2, []int64{2, 3}, 6},
	}
	for _, c := range testCases {
		//when
		gotFactors, gotProduct := rowRight(givenGrid, c.count)
		//then
		if !commonintutils.Int64SliceEquals(gotFactors, c.wantFactors) {
			t.Errorf("rowRight(%v, %v) == %v, %v want %v, %v",
				givenGrid,
				c.count,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
		//and
		if gotProduct != c.wantProduct {
			t.Errorf("rowRight(%v, %v) == %v, %v want %v, %v",
				givenGrid,
				c.count,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
	}
}

func TestColumnDown(t *testing.T) {

	//given
	givenGrid := [][]int64{
		[]int64{1, 1, 1, 1},
		[]int64{1, 2, 1, 1},
		[]int64{1, 3, 1, 1},
		[]int64{1, 1, 1, 1},
	}

	testCases := []struct {
		count       int
		wantFactors []int64
		wantProduct int64
	}{
		//where
		{0, []int64{}, 0},
		{2, []int64{2, 3}, 6},
	}
	for _, c := range testCases {
		//when
		gotFactors, gotProduct := columnDown(givenGrid, c.count)
		//then
		if !commonintutils.Int64SliceEquals(gotFactors, c.wantFactors) {
			t.Errorf("columnDown(%v, %v) == %v, %v want %v, %v",
				givenGrid,
				c.count,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
		//and
		if gotProduct != c.wantProduct {
			t.Errorf("columnDown(%v, %v) == %v, %v want %v, %v",
				givenGrid,
				c.count,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
	}
}

func TestGetLargestProductInSlice(t *testing.T) {
	//given
	testCases := []struct {
		slice       []int64
		size        int
		wantFactors []int64
		wantProduct int64
	}{
		//where
		{[]int64{}, 0, []int64{}, 0},
		{[]int64{1, 2, 3, 4}, 2, []int64{3, 4}, 12},
		{[]int64{1, 3, 4, 2}, 2, []int64{3, 4}, 12},
		{[]int64{1, 3, 4, 2, 5}, 3, []int64{4, 2, 5}, 40},
	}
	for _, c := range testCases {
		//when
		gotFactors, gotProduct := getLargestProductInSlice(c.slice, c.size)
		//then
		if !commonintutils.Int64SliceEquals(gotFactors, c.wantFactors) {
			t.Errorf("getLargestProductInSlice(%v, %v) == %v, %v want %v, %v",
				c.slice,
				c.size,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
		//and
		if gotProduct != c.wantProduct {
			t.Errorf("getLargestProductInSlice(%v, %v) == %v, %v want %v, %v",
				c.slice,
				c.size,
				gotFactors,
				gotProduct,
				c.wantFactors,
				c.wantProduct,
			)
		}
	}
}

func TestCreateColumnSlice(t *testing.T) {
	//given
	testCases := []struct {
		grid   [][]int64
		column int
		want   []int64
	}{
		//where
		{[][]int64{}, 0, []int64{}},
		{[][]int64{
			[]int64{1},
		}, 0, []int64{1}},
		{[][]int64{
			[]int64{1, 2},
			[]int64{3, 4},
		}, 0, []int64{1, 3}},
		{[][]int64{
			[]int64{1, 2},
			[]int64{3, 4},
		}, 1, []int64{2, 4}},
	}
	for _, c := range testCases {
		//when
		got := createColumnSlice(c.grid, c.column)
		//then
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("createColumnSlice(%v, %v) == %v, want %v", c.grid, c.column, got, c.want)
		}
	}
}

func TestGetGridSize(t *testing.T) {
	//given
	testCases := []struct {
		grid        [][]int64
		wantRows    int
		wantColumns int
	}{
		//where
		{[][]int64{}, 0, 0},
		{[][]int64{
			[]int64{1},
		}, 1, 1},
		{[][]int64{
			[]int64{1, 2},
			[]int64{3, 4},
		}, 2, 2},
		{[][]int64{
			[]int64{1, 2, 3},
			[]int64{4, 5, 6},
		}, 2, 3},
	}
	for _, c := range testCases {
		//when
		gotRows, gotColumns := getGridSize(c.grid)
		//then
		if gotRows != c.wantRows {
			t.Errorf("getGridSize(%v) == %v, want %v", c.grid, gotRows, c.wantRows)
		}
		//and
		if gotColumns != c.wantColumns {
			t.Errorf("getGridSize(%v) == %v, want %v", c.grid, gotColumns, c.wantColumns)
		}
	}
}

func TestGenerateGrid(t *testing.T) {
	//given
	testCases := []struct {
		wantColumn int
		wantRow    int
	}{
		//where
		{20, 20}, //provided example
	}

	for _, c := range testCases {
		//when
		got := generateGrid()
		//then
		if !checkGrid(c.wantColumn, c.wantRow, got) {
			t.Errorf("generateGrid() == %v, want %v, %v", got, c.wantColumn, c.wantRow)
		}
	}
}

// test helper function
func checkGrid(columns int, rows int, grid [][]int64) bool {
	if len(grid) != rows {
		fmt.Println(fmt.Sprintf("expected %v rows, got %v", rows, len(grid)))
		return false
	}
	for _, row := range grid {
		if len(row) != columns {
			fmt.Println(fmt.Sprintf("expected %v columns, got %v", columns, len(row)))
			return false
		}
	}
	return true
}

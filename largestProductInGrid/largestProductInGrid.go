package main

/*
In the 20×20 grid below, four numbers along a diagonal line have been marked with a star.

08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 *26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 *63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 *78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 *14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48

The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

What is the greatest product of four adjacent numbers in the same direction
(up, down, left, right, diagonal left, diagonal right) in the 20×20 grid?
*/

import (
	"fmt"
	"projectEuler/common/commonintutils"
)

// Cell - a representation of a location in a grid
type Cell struct {
	row    int
	column int
}

func main() {
	count := 4

	grid := generateGrid()

	fmt.Println(fmt.Sprintf("Given 20X20 grid: %v", grid))
	fmt.Println(fmt.Sprintf("Finding largest product of %v sequential numbers:", count))
	factors, product := columnDown(grid, count)
	fmt.Println(fmt.Sprintf("column down/up %v=%v:", factors, product))
	factors, product = rowRight(grid, count)
	fmt.Println(fmt.Sprintf("row right/left: %v=%v", factors, product))
	factors, product = slash(grid, count, "left")
	fmt.Println(fmt.Sprintf("slash left: %v=%v", factors, product))
	factors, product = slash(grid, count, "right")
	fmt.Println(fmt.Sprintf("slash right: %v=%v", factors, product))
}

//up will be the same product
func columnDown(grid [][]int64, count int) (factors []int64, largestProduct int64) {
	_, columns := getGridSize(grid)

	//step through the count of columns
	for column := 0; column < columns; column++ {
		//create slice of all column values out of the rows
		columnSlice := createColumnSlice(grid, column)
		//chunk through the slice to find the largest product of n factors
		f, p := getLargestProductInSlice(columnSlice, count)
		//if larger than current then assign
		if p > largestProduct {
			factors = f
			largestProduct = p
		}
	}
	//return the factors and the product
	return
}

//left will be the same product
func rowRight(grid [][]int64, count int) (factors []int64, largestProduct int64) {
	//step through the rows
	for _, row := range grid {
		//chunk through the row to find the largest product of n factors
		f, p := getLargestProductInSlice(row, count)
		//if larger than current then assign
		if p > largestProduct {
			factors = f
			largestProduct = p
		}
	}
	//return the factors and the product
	return
}

func slash(grid [][]int64, count int, direction string) (factors []int64, largestProduct int64) {
	rows, columns := getGridSize(grid)

	if rows == 0 {
		return []int64{}, 0
	}

	rowEnd, columnEnd := getEndValues(rows, columns, direction)

	rowStart, columnStart := getStartValues(rows, columns, direction)
	diagonal := []Cell{Cell{rowStart, columnStart}}
	//increment the diagonal each time until you are on the last diagonal
	for !(diagonal[0].row == rowEnd && diagonal[0].column == columnEnd) {
		//create slice of all slash right values
		ds := getDiagonalSlice(diagonal, grid)
		if len(ds) >= count {
			f, p := getLargestProductInSlice(ds, count)

			// compare and save
			if p > largestProduct {
				factors = f
				largestProduct = p
			}
		}

		diagonal = incrementDiagonal(diagonal, rows, columns, direction)
	}
	return
}

func incrementDiagonal(diagonal []Cell, rows int, columns int, direction string) []Cell {
	switch direction {
	case "right":
		return nextRightDiagonal(diagonal, rows, columns)
	case "left":
		return nextLeftDiagonal(diagonal, rows, columns)
	default:
		return []Cell{}
	}
}

func getEndValues(rows int, columns int, direction string) (endRow int, endColumn int) {
	switch direction {
	case "right":
		return rows - 1, columns - 1
	case "left":
		return rows - 1, 0
	default:
		return rows - 1, columns - 1
	}
}

func getStartValues(rows int, columns int, direction string) (endRow int, endColumn int) {
	switch direction {
	case "right":
		return 0, 0
	case "left":
		return 0, columns - 1
	default:
		return 0, 0
	}
}

func nextLeftDiagonal(diagonalIn []Cell, rows int, columns int) (newDiagonal []Cell) {
	if len(diagonalIn) == 0 {
		return []Cell{Cell{0, columns - 1}}
	}

	//increment all existing rows to next column remove any that are over the end
	newDiagonal = []Cell{}
	for _, cell := range diagonalIn {
		row := cell.row
		column := cell.column
		if column > 0 {
			newDiagonal = append(newDiagonal, Cell{row, column - 1})
		}
	}

	//if not bottom add new row
	size := getMaxRow(newDiagonal)
	if size+1 < rows {
		newDiagonal = append(newDiagonal, Cell{size + 1, columns - 1})
	}

	return
}

func nextRightDiagonal(diagonalIn []Cell, rows int, columns int) (newDiagonal []Cell) {
	if len(diagonalIn) == 0 {
		return []Cell{Cell{0, 0}}
	}

	//increment all existing rows to next column remove any that are over the end
	newDiagonal = []Cell{}
	for _, cell := range diagonalIn {
		row := cell.row
		column := cell.column
		if column+1 < columns {
			newDiagonal = append(newDiagonal, Cell{row, column + 1})
		}
	}

	//if not bottom add new row
	size := getMaxRow(newDiagonal)
	if size+1 < rows {
		newDiagonal = append(newDiagonal, Cell{size + 1, 0})
	}

	return
}

func getMaxRow(diagonal []Cell) (maxRow int) {
	for _, cell := range diagonal {
		if cell.row > maxRow {
			maxRow = cell.row
		}
	}
	return
}

func getDiagonalSlice(diagonal []Cell, grid [][]int64) (diagonalSlice []int64) {
	for _, cell := range diagonal {
		diagonalSlice = append(diagonalSlice, grid[cell.row][cell.column])
	}
	return
}

func generateGrid() [][]int64 {
	return [][]int64{
		[]int64{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
		[]int64{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
		[]int64{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
		[]int64{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91},
		[]int64{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		[]int64{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		[]int64{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		[]int64{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		[]int64{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		[]int64{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
		[]int64{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92},
		[]int64{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
		[]int64{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		[]int64{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40},
		[]int64{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		[]int64{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		[]int64{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		[]int64{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16},
		[]int64{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54},
		[]int64{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48},
	}
}

func getLargestProductInSlice(sl []int64, count int) (factors []int64, largestProduct int64) {
	for index := range sl {
		if index+count <= len(sl) {
			f := sl[index : index+count]
			p := commonintutils.ProductInt64Slice(f)
			if p > largestProduct {
				factors = f
				largestProduct = p
			}
		}
	}
	return
}

func createColumnSlice(grid [][]int64, column int) (columnSlice []int64) {
	for _, row := range grid {
		columnSlice = append(columnSlice, row[column])
	}
	return
}

func getGridSize(grid [][]int64) (rows int, columns int) {
	if len(grid) == 0 {
		return 0, 0
	}
	return len(grid), len(grid[0])
}

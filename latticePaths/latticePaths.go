package main

/*
* Starting in the top left corner of a 2×2 grid, and only being able to move to the
* right and down, there are exactly 6 routes to the bottom right corner.
* How many such routes are there through a 20×20 grid?
 */

import (
	"fmt"
)

func main() {
	//gridSize := 1 //1X1 = 2 paths
	//gridSize := 2 //2X2 = 6 paths
	//gridSize := 3 //3X3 = 20 paths
	gridSize := 20 //20X20 = 137846528820 paths
	//gridSize := 70 //70X70 = 17130505608995310632 paths
	//gridSize := 71 //71X71 - Exceeds max uint64 2^64 − 1 which equals 18,446,744,073,709,551,615
	fmt.Println(fmt.Sprintf("gridSize: [%vx%v]", gridSize, gridSize))

	easyAnswer := easyWay(gridSize)
	fmt.Println(fmt.Sprintf("easyAnswer: %v paths", easyAnswer))
}

func easyWay(gridSize int) uint64 {
	//convert to 1 based before making the calls
	//TODO: convert back to 0 based
	//TODO: the right/down coordinates are reversed, not a big issue, but confusing
	//TODO: convert to the golang Big library
	gridSize++

	myGrid := initGrid(gridSize)
	fillGrid(gridSize, gridSize, gridSize, myGrid)

	//fmt.Println(fmt.Sprintf("filled grid: %v", myGrid))
	return (*myGrid)[0][0]
}

func fillGrid(gridRight int, gridDown int, gridSize int, myGrid *[][]uint64) {
	//because I'm using a uint64 I have to trigger on 0, i.e. locations are 1 based
	if gridRight > 0 && gridDown > 0 {
		checkRight(gridRight, gridDown, myGrid)
		checkDown(gridRight, gridDown, myGrid)
		// locations are 1 based
		if gridDown == 1 {
			fillGrid(gridRight-1, gridSize, gridSize, myGrid)
		}
		fillGrid(gridRight, gridDown-1, gridSize, myGrid)
	}
}

func checkRight(gridRight int, gridDown int, myGrid *[][]uint64) {
	if gridRight < len((*myGrid)) {
		(*myGrid)[gridRight-1][gridDown-1] += (*myGrid)[gridRight][gridDown-1]
	}
}

func checkDown(gridRight int, gridDown int, myGrid *[][]uint64) {
	if gridDown < len((*myGrid)) {
		(*myGrid)[gridRight-1][gridDown-1] += (*myGrid)[gridRight-1][gridDown]
	}
}

func initGrid(gridSize int) *[][]uint64 {
	myGrid := make([][]uint64, gridSize)
	for i := range myGrid {
		myGrid[i] = make([]uint64, gridSize)
	}
	myGrid[gridSize-1][gridSize-1] = 1
	return &myGrid
}

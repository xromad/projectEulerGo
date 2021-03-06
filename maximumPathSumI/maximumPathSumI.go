package main

/*
By starting at the top of the triangle below and moving to adjacent numbers
on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                           75
                          95 64
                        17 47 82
                      18 35 87 10
                    20 04 82 47 65
                  19 01 23 75 03 34
                88 02 77 73 07 63 67
              99 65 04 28 06 16 70 92
            41 41 26 56 83 40 80 70 33
          41 48 72 33 47 32 37 16 94 29
        53 71 44 65 25 43 91 52 97 51 14
      70 11 33 28 77 73 17 78 39 68 17 57
    91 71 52 38 17 14 91 43 58 50 27 29 48
  63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve
this problem by trying every route. However, Problem 67,
is the same challenge with a triangle containing one-hundred
rows; it cannot be solved by brute force, and requires a clever
method! ;o)
*/

import (
	"fmt"
)

type node struct {
	value uint
	total uint
	left  *node
	right *node
}

func main() {
	//sample := true //provided sample
	sample := false
	tree := makeTree(sample)
	fmt.Println(fmt.Sprintf("calculating the maximum path sum for the tree starting with node: %v", *tree))
	fmt.Println(fmt.Sprintf("maximum path sum: %v", checkTree(tree)))
	fmt.Println(fmt.Sprintf("values for nodes in path: %v", getNodes(tree)))
}

func getNodes(n *node) (path []uint) {
	if n.left != nil {
		if n.left.total > n.right.total {
			path = append(getNodes(n.left), n.value)
		} else {
			path = append(getNodes(n.right), n.value)
		}
	} else {
		path = append(path, n.value)
	}
	return
}

func checkTree(activeNode *node) uint {
	if activeNode.total > 0 {
		return activeNode.total
	}

	var leftTotal, rightTotal uint
	if activeNode.left != nil {
		leftTotal = checkTree(activeNode.left)
	} else {
		activeNode.total = activeNode.value
		return activeNode.total
	}

	if activeNode.right != nil {
		rightTotal = checkTree(activeNode.right)
	}

	if leftTotal > rightTotal {
		activeNode.total = leftTotal + activeNode.value
	} else {
		activeNode.total = rightTotal + activeNode.value
	}

	return activeNode.total
}

func makeTree(sample bool) *node {
	if sample {
		return makeExampleTree()
	}
	return makeRealTree()
}

func makeExampleTree() (tree *node) { //the provided example
	// build it upside down
	row4 := buildRow([]uint{8, 5, 9, 3}, true, nil)
	row3 := buildRow([]uint{2, 4, 6}, false, row4)
	row2 := buildRow([]uint{7, 4}, false, row3)
	row1 := buildRow([]uint{3}, false, row2)
	tree = &row1[0]
	return
}

//always assumes the previous list is larger than the current list
func buildRow(values []uint, first bool, previous []node) (nodes []node) {
	left := 0
	right := 1
	for _, value := range values {
		if first {
			nodes = append(nodes, node{value, 0, nil, nil})
		} else {
			nodes = append(nodes, node{value, 0, &previous[left], &previous[right]})
			left++
			right++
		}
	}
	return
}

func makeRealTree() (tree *node) {
	// build it upside down
	row15 := buildRow([]uint{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}, true, nil)
	row14 := buildRow([]uint{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}, false, row15)
	row13 := buildRow([]uint{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}, false, row14)
	row12 := buildRow([]uint{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}, false, row13)
	row11 := buildRow([]uint{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}, false, row12)
	row10 := buildRow([]uint{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}, false, row11)
	row9 := buildRow([]uint{41, 41, 26, 56, 83, 40, 80, 70, 33}, false, row10)
	row8 := buildRow([]uint{99, 65, 04, 28, 06, 16, 70, 92}, false, row9)
	row7 := buildRow([]uint{88, 02, 77, 73, 07, 63, 67}, false, row8)
	row6 := buildRow([]uint{19, 01, 23, 75, 03, 34}, false, row7)
	row5 := buildRow([]uint{20, 04, 82, 47, 65}, false, row6)
	row4 := buildRow([]uint{18, 35, 87, 10}, false, row5)
	row3 := buildRow([]uint{17, 47, 82}, false, row4)
	row2 := buildRow([]uint{95, 64}, false, row3)
	row1 := buildRow([]uint{75}, false, row2)
	tree = &row1[0]
	return
}

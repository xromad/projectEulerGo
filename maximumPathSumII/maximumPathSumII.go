package main

/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below,
the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.
NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value uint
	total uint
	left  *node
	right *node
}

func main() {
	//tree, rows := makeTree(true) //provided sample
	tree, rows := makeTree(false) //the problem

	fmt.Println(fmt.Sprintf("calculating the maximum path sum for a tree with %v rows", rows))
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

func readFromFile(filename string) (contents [][]string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func makeTree(sample bool) (*node, uint) {
	if sample {
		return parseTree("sampleTriangle.txt")
	}
	return parseTree("p067_triangle.txt")
}

func parseTree(fileName string) (tree *node, rows uint) {
	content := readFromFile(fileName)
	rows = uint(len(content))

	previousLine := []*node{}
	for line := len(content) - 1; line >= 0; line-- {
		currentLine := []*node{}
		for index, value := range content[line] {
			num, _ := strconv.Atoi(value)
			if len(previousLine) == 0 {
				currentLine = append(currentLine, &node{uint(num), 0, nil, nil})
			} else {
				currentLine = append(currentLine, &node{uint(num), 0, previousLine[index], previousLine[index+1]})
			}
		}
		previousLine = currentLine
		if len(currentLine) == 1 {
			tree = currentLine[0]
		}
	}
	return
}

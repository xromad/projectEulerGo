package main

/*
The following iterative sequence is defined for the set of positive integers:
n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
NOTE: Once the chain starts the terms are allowed to go above one million.
*/

import (
	"fmt"
)

func main() {
	//var max int64 = 13 // the provided example
	var max int64 = 1000000 // 1 million
	fmt.Println(fmt.Sprintf("finding longest Colatz Sequence with a start below: %v", max))

	longest := findLongest(max)
	fmt.Println(fmt.Sprintf("longest: %v", longest))
}

func findLongest(max int64) (longest []int64) {
	for i := int64(1); i <= max; i++ {
		candidate := generateSequence(i)
		if len(candidate) > len(longest) {
			longest = candidate
		}
	}
	return longest
}

func generateSequence(num int64) (sequence []int64) {
	sequence = append(sequence, num)
L:
	for {
		switch {
		case num == 1:
			break L
		case num%2 == 0:
			num = evenFunction(num)
			sequence = append(sequence, num)
		default:
			num = oddFunction(num)
			sequence = append(sequence, num)
		}
	}
	return sequence
}

func evenFunction(num int64) int64 {
	return num / 2
}

func oddFunction(num int64) int64 {
	return num*3 + 1
}

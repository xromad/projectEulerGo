package main

import (
	"fmt"
	"math/big"
)

/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

type fibonacciNumber struct {
	num   big.Int
	index int
}

func main() {
	getIndex(1000)
}

func getIndex(digits int) fibonacciNumber {
	fmt.Println(fmt.Sprintf("Finding the index for the first %v digit Fibonacci number", digits))

	previousFibNum := fibonacciNumber{*big.NewInt(int64(1)), 1}
	fibNum := fibonacciNumber{*big.NewInt(int64(1)), 2}
	for len(fibNum.num.String()) < digits {
		fibNum, previousFibNum = getNextFibonacciNumber(previousFibNum, fibNum)
	}

	fmt.Println(fmt.Sprintf("The first %v digit Fibonacci number is: %v, index: %v", digits, fibNum.num, fibNum.index))
	return fibNum
}

func getNextFibonacciNumber(prevFibIn fibonacciNumber, fibIn fibonacciNumber) (fibonacciNumber, fibonacciNumber) {
	nextFib := *prevFibIn.num.Add(&prevFibIn.num, &fibIn.num)
	return fibonacciNumber{nextFib, fibIn.index + 1}, fibIn
}

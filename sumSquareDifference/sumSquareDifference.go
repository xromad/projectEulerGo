package main

/*
 * The sum of the squares of the first ten natural numbers is,
 * 1^2 + 2^2 + ... + 10^2 = 385
 *
 * The square of the sum of the first ten natural numbers is,
 * (1 + 2 + ... + 10)^2 = 55^2 = 3025
 *
 * Hence the difference between the sum of the squares of the first ten natural numbers and the
 * square of the sum is 3025 − 385 = 2640.
 *
 * Find the difference between the sum of the squares of the first one hundred natural numbers
 * and the square of the sum.
 */

import (
	"fmt"
)

func main() {
	var min int64 = 1
	//var max int64 = 10 // 2640
	//var max int64 = 20 // 41230
	var max int64 = 100 // 25164150

	fmt.Println(fmt.Sprintf("sum square difference of numbers between: %v and %v", min, max))
	sumOfSquares, squareOfSum, difference := findSumSquareDifference(min, max)
	fmt.Println(fmt.Sprintf("difference between sumOfSquares: %v and squareOfSum %v = %v",
		sumOfSquares,
		squareOfSum,
		difference,
	))
}

func findSumSquareDifference(min int64, max int64) (sumOfSquares int64, squareOfSum int64, difference int64) {

	sumOfSquares = getSumOfSquares(min, max)
	squareOfSum = getSquareOfSum(min, max)
	return sumOfSquares, squareOfSum, squareOfSum - sumOfSquares
}

func getSumOfSquares(min int64, max int64) (result int64) {
	for i := min; i <= max; i++ {
		result = result + expInt64(i, 2)
	}
	return
}

func getSquareOfSum(min int64, max int64) (result int64) {
	var sum int64
	for i := min; i <= max; i++ {
		sum = sum + i
	}
	return expInt64(sum, 2)
}

func expInt64(x int64, n int64) (result int64) {
	if n == 0 {
		return 1
	}
	result = x
	for i := int64(1); i < n; i++ {
		result = result * x
	}
	return
}

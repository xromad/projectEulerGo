package main

import (
	"fmt"
	"projectEuler/common/commonintutils"
)

/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

func main() {
	getAmicableNumbersSum(10000)
}

func getAmicableNumbersSum(largestNumber int64) (sum int64) {
	fmt.Println(fmt.Sprintf("Calculating Amicable Numbers Sum For: %v", largestNumber))
	amicableNumbers := getAmicableNumbers(largestNumber)
	sum = commonintutils.SumInt64Slice(amicableNumbers)
	fmt.Println(fmt.Sprintf("Factoral Digit Sum For: %v = %v", largestNumber, sum))
	return
}

func getAmicableNumbers(n int64) (numbers []int64) {
	numbersMap := make(map[int64]int64)
	for i := int64(1); i <= n; i++ {
		sumOfProperDivisors := commonintutils.SumInt64Slice(getProperDivisors(i))
		numbersMap[i] = sumOfProperDivisors

		if numbersMap[sumOfProperDivisors] != 0 && numbersMap[sumOfProperDivisors] == i && i != sumOfProperDivisors {
			numbers = append(numbers, sumOfProperDivisors)
			numbers = append(numbers, i)
		}
	}
	return
}

func getProperDivisors(n int64) (divisors []int64) {
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return
}

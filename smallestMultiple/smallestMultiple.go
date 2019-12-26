package main

/*
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 * Note: I originally took a prime factorization route that was painfully slow
 * this way is much much faster
 */

import (
	"fmt"
	"projectEuler/common/commonintutils"
)

func main() {
	findSmallestMultiple(1, 20)
}

func findSmallestMultiple(min int64, max int64) (product int64) {
	fmt.Println(fmt.Sprintf("Finding smallest mulitple of all numbers between: %v and %v", min, max))
	product = 1
	maxProduct := commonintutils.ProductInt64Range(min, max)

	for ; product <= maxProduct; product++ {
		match := true
		for i := min; i <= max; i++ {
			if product%i != 0 {
				match = false
				break
			}
		}

		if match {
			fmt.Println(fmt.Sprintf("Smallest multiple of factors from: %v to %v=%v", min, max, product))
			return product
		}
	}
	fmt.Println(fmt.Sprintf("Smallest multiple of factors from: %v to %v not found", min, max))
	return 0
}

package main

/*
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 * I originally took a prime factorization route that was painfully slow
 * this way is much much faster
 */

import (
	"fmt"
	"projectEuler/common/commonintutils"
)

func main() {
	var min int64 = 1
	//var max int64 = 10 // 2520
	//var max int64 = 11 // 27720
	var max int64 = 20 // 232792560

	fmt.Println(fmt.Sprintf("Finding smallest mulitple of all numbers between: %v and %v", min, max))
	product := findSmallestMultiple(min, max)
	fmt.Println(fmt.Sprintf("Smallest multiple of factors from: %v to %v=%v", min, max, product))
}

func findSmallestMultiple(min int64, max int64) (product int64) {
	product = 1
	maxProduct := commonintutils.ProductInt64Range(min, max)

	for ; product <= maxProduct; product++ {
		if product%100000 == 0 {
			fmt.Print(".")
		}

		match := true
		for i := min; i <= max; i++ {
			if product%i != 0 {
				match = false
				break
			}
		}

		if match {
			return product
		}
	}
	return 0
}

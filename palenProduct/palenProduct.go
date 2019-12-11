package main

/*
 * A palindromic number reads the same both ways.
 * The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

import (
	"projectEuler/common/commonintutils"
	"projectEuler/common/commonstringutils"
	"fmt"
)

func main() {
	// 100 is the smallest 3 digit number
	var min int64 = 100
	// 999 is the largest 3 digit number
	var max int64 = 999

	fmt.Println(fmt.Sprintf("Finding palendromic product of two numberes between: %v and %v",
		min,
		max,
	))

	first, second, product := findLargestProduct(min, max)
	fmt.Println(fmt.Sprintf("Largest palendrome product: %v*%v=%v", first, second, product))
}

func findLargestProduct(min int64, max int64) (lowestFactor int64, highestFactor int64, largestPalendrome int64) {
	lowestFactor = min
	highestFactor = max

	for secondFactor := max; secondFactor >= lowestFactor; secondFactor-- {
		for firstFactor := max; firstFactor >= lowestFactor; firstFactor-- {
			product := firstFactor * secondFactor
			if commonstringutils.IsPalendrome(commonintutils.I64toA(product)) {
				if product > largestPalendrome {
					largestPalendrome = product
					highestFactor = secondFactor
					lowestFactor = firstFactor
					break
				}
			}
		}
	}
	return
}

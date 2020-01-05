package main

import (
	"fmt"
	"projectEuler/common/commonintutils"
)

/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means
that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is
called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
the smallest number that can be written as the sum of two abundant numbers is 24.

By mathematical analysis, it can be shown that all integers greater than 28123 can be written
as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis
even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is
less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

type numType int

const (
	deficient    numType = -1
	perfect              = 0
	abundant             = 1
	numTypeError         = 2
)

func main() {
	getNonAbundentSums(28123)
}

func getNonAbundentSums(max int64) (sum int64) {
	fmt.Println(fmt.Sprintf("finding sum of nonAbundent integers below %v not expressed as two abundent numbers", max))

	var abundantNumbers []int64
	var nonAbundantNumbers []int64

	for i := int64(1); i <= max; i++ {
		numberType := getNumberType(i)
		if !isSumOfabundants(abundantNumbers, i) {
			nonAbundantNumbers = append(nonAbundantNumbers, i)
		}
		if numberType == abundant {
			abundantNumbers = append(abundantNumbers, i)
		}
	}
	sum = commonintutils.SumInt64Slice(nonAbundantNumbers)
	fmt.Println(fmt.Sprintf("Answer = %v", sum))
	return
}

func isSumOfabundants(abundantNumbers []int64, number int64) bool {
	for lower := 0; lower < len(abundantNumbers); lower++ {
		for upper := lower; upper < len(abundantNumbers); upper++ {
			if abundantNumbers[lower]+abundantNumbers[upper] == number {
				return true
			}
		}
	}
	return false
}

func getNumberType(number int64) (numberType numType) {
	sumofFactors := commonintutils.SumInt64Slice(getFactors(number))
	switch {
	case sumofFactors < number:
		return deficient
	case sumofFactors == number:
		return perfect
	case sumofFactors > number:
		return abundant
	}
	fmt.Println(fmt.Sprintf("Error: Evaluating %v resulted in a numTypeError", number))
	return numTypeError
}

func getFactors(number int64) (factors []int64) {
	if number == 1 {
		return []int64{1}
	}

	for i := int64(1); i <= number/2; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return
}

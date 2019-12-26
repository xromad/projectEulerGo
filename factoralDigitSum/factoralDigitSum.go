package main

import (
	"fmt"
	"projectEuler/common/commonintutils"
	"strconv"
)

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func main() {
	factoralDigitSum(100)
	//drat 100! blows an int64 have to go to big
}

func factoralDigitSum(n int64) (sum int64) {
	fmt.Println(fmt.Sprintf("Calculating Factoral Digit Sum For: %v", n))
	nFact := commonintutils.BigFactoral(n)
	fmt.Println(fmt.Sprintf("n Factoral: %v", nFact.String()))
	for _, r := range nFact.String() {
		d, e := strconv.Atoi(string(r))
		if e != nil {
			fmt.Println(fmt.Sprintf("Error: %s", e))
		}
		digit := int64(d)
		sum += digit
	}
	fmt.Println(fmt.Sprintf("Factoral Digit Sum is %v", sum))
	return sum
}

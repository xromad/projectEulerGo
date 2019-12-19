package main

import (
	"fmt"
	"math"
	"projectEuler/common/commonintutils"
)

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the values for a, b, c.
*/

func main() {
	//var targetSum int64 = 12
	//var targetSum int64 = 30
	var targetSum int64 = 1000

	fmt.Println(fmt.Sprintf("Finding Pythagorean triplet where a + b + c = %v", targetSum))
	digitsSlice := getPythagTripletEqN(targetSum)
	if commonintutils.SumInt64Slice(digitsSlice) != 0 {
		sumOfDigits := commonintutils.SumInt64Slice(digitsSlice)
		fmt.Println(fmt.Sprintf(
			"%v^2+%v^2=%v^2",
			digitsSlice[0],
			digitsSlice[1],
			digitsSlice[2],
		))
		fmt.Println(fmt.Sprintf(
			"%v+%v+%v=%v",
			digitsSlice[0],
			digitsSlice[1],
			digitsSlice[2],
			sumOfDigits,
		))
	} else {
		fmt.Println("Not found!")
	}
}

func getPythagTripletEqN(targetSum int64) (digitSlice []int64) {
	digitSlice = []int64{1, 1, 1}

	for {
		switch {
		//if a is greater than target then no need to continue
		case digitSlice[0] == 0:
			return []int64{0, 0, 0}
		//found our solution
		case commonintutils.SumInt64Slice(digitSlice) == targetSum:
			return digitSlice
		//keep looking
		default:
			digitSlice = findNextPythagTriplet(
				digitSlice[0],
				digitSlice[1]+1,
				targetSum)
		}
	}
}

func findNextPythagTriplet(aIn int64, bIn int64, max int64) []int64 {
	stop := false
	bStart := bIn
	for a := aIn; !stop && a <= max; a++ {
		for b := bStart; !stop && b <= max; b++ {
			if b > a {
				csq := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
				f := math.Sqrt(csq)
				if f == math.Floor(f) {
					return []int64{a, b, int64(f)}
				}
			}
		}
		bStart = 1
	}
	return []int64{0, 0, 0}
}

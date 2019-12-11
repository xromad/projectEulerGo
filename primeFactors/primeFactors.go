package main

/*
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 600851475143 ?
*/

import (
	"projectEuler/common/commonintutils"
	"projectEuler/common/primesiev"
	"fmt"
	"strings"
)

func main() {
	//var numToFactor int64 = 2520
	var numToFactor int64 = 600851475143
	//var numToFactor int64 = 600

	fmt.Println(fmt.Sprintf("Finding prime factors of %v:", numToFactor))

	factors := factorNumber(numToFactor)
	fmt.Println(fmt.Sprintf("factors =: %v", factors))
	fmt.Println(fmt.Sprintf("equation =: %s", genEquation(factors)))
	fmt.Println(fmt.Sprintf("max prime factor: %v", factors[len(factors)-1]))
}

func genEquation(factors []int64) string {
	if len(factors) == 0 {
		return "0=0"
	}
	return fmt.Sprintf("%v=%v", strings.Join(commonintutils.Int64SliceToStringSlice(factors), "*"),
		commonintutils.ProductInt64Slice(factors))
}

func factorNumber(number int64) (factors []int64) {
	halfInitialNumber := number / 2
	primes := primesiev.GetNextPrime(&[]int64{})
	factors = append(factors, 1)

	if number == 0 || number == 1 {
		return []int64{1}
	}

	for factored := false; !factored; {
		for _, p := range primes {
			if p != 1 {
				if number >= p && number%p == 0 {
					factors = append(factors, p)
					number = number / p
					break
				}
			}
			if number == 1 || p > halfInitialNumber {
				if number > 1 {
					factors = append(factors, number)
				}
				factored = true
				break
			}
			primes = primesiev.GetNextPrime(&primes)
		}
	}
	return
}

package main

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/

import (
	"fmt"
	"projectEuler/common/primesiev"
)

func main() {
	getSumOfPrimes(2000000)
}

func getSumOfPrimes(max int64) (sumOfPrimes int64) {
	fmt.Println(fmt.Sprintf("Finding the sum of primes below %v:", max))
	if max == 0 {
		return 0
	}

	primes := []int64{2}

	for prime := int64(2); prime <= max; {
		sumOfPrimes = sumOfPrimes + prime
		primes := primesiev.GetNextPrime(&primes)
		prime = primes[len(primes)-1]
	}

	fmt.Println(fmt.Sprintf("sum of primes = %v", sumOfPrimes))
	return sumOfPrimes
}

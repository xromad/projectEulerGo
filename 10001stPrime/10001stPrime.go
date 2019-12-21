package main

/*
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10,001st prime number?
 */

import (
	"fmt"
	"projectEuler/common/primesiev"
)

func main() {
	findNthPrime(10001)
}

func findNthPrime(primeCount int) []int64 {
	fmt.Println(fmt.Sprintf("Finding %dst prime:", primeCount))

	primes := primesiev.FindPrimesFrom(primeCount, 2)
	fmt.Println(fmt.Sprintf("prime #%d is %v", primeCount, primes[len(primes)-1]))
	return primes
}

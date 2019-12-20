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
	primeCount := 10001
	fmt.Println(fmt.Sprintf("Finding %dst prime:", primeCount))

	primes := findNPrimes(primeCount)
	fmt.Println(fmt.Sprintf("Largest %dst prime is %v", primeCount, primes[len(primes)-1]))
}

func findNPrimes(primeCount int) []int64 {
	return primesiev.FindPrimesFrom(primeCount, 2)
}

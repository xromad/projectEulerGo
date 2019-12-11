// Package primesiev has functionality to produce a slice of primes
package primesiev

// GetNextPrime - returns the next prime in a slice
func GetNextPrime(primes *[]int64) []int64 {
	if len(*primes) == 0 {
		*primes = []int64{1}
		return *primes
	}

	for i, found := (*primes)[int64(len(*primes)-1)]+1, false; !found; i++ {
		if isPrime(i, *primes) {
			*primes = append(*primes, i)
			found = true
		}
	}

	return *primes
}

// FindPrimes - returns the first n primes in a slice
func FindPrimes(findCount int) (primes []int64) {
	for i := int64(1); len(primes) < findCount; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return
}

// FindPrimesFrom - returns the first n primes in a slice starting at m
func FindPrimesFrom(findCount int, start int) (primes []int64) {
	if findCount == 0 {
		return []int64{}
	}
	if findCount == 1 && start == 0 {
		return []int64{1}
	}
	for i := int64(start); len(primes) < findCount; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return
}

// FindPrimesSmallerThan - returns a slce of primes smaller than the set number
func FindPrimesSmallerThan(limit int64) (primes []int64) {
	primes = append(primes, 1)
	for i := int64(2); i <= limit; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return
}

// TODO: refactor - only have to search through primes smaller than 1/2 n
func isPrime(n int64, primes []int64) bool {
	if n == 1 {
		return true
	}
	for _, p := range primes {
		if p != 1 && (n%p) == 0 {
			return false
		}
	}
	return true
}

// PrimeFactorInt64 - take an int64 number and return the prime factors
func PrimeFactorInt64(number int64) (factors []int64) {
	if number == 0 || number == 1 {
		return []int64{1}
	}

	halfInitialNumber := number / 2
	primes := GetNextPrime(&[]int64{})
	factors = append(factors, 1)

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
			primes = GetNextPrime(&primes)
		}
	}
	return
}

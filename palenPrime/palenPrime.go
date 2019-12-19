package main

/*
 * Look for two types of reversable primes: Palendrome and Emirp
 * Palendrome primes are like "101" where reversed it is the same prime
 * Emirp primes are like "13:31" where the reverse is another different prime
 *  - palendromes are 'kind of emirp' but we specifically exclude them from the definition
 */

import (
	"fmt"
	"projectEuler/common/commonstringutils"
	"projectEuler/common/primesiev"
	"strconv"
)

func main() {
	findCount := 2000
	fmt.Println(fmt.Sprintf("Finding first %d primes:", findCount))

	primes := primesiev.FindPrimes(findCount)
	fmt.Println(fmt.Sprintf("Largest prime: %v", primes[len(primes)-1]))
	fmt.Println(fmt.Sprintf("Palendromic primes: %v", findPalendromes(primes)))
	fmt.Println(fmt.Sprintf("Emirp primes: %v", findEmirps(primes)))
}

func findEmirps(primes []int64) (retMap map[int64]int64) {
	//note: have to make the map or you get a panic
	retMap = make(map[int64]int64)
	for _, p := range primes {
		lower, higher, isEmirp := checkEmirp(primes, p)
		if isEmirp {
			if retMap[lower] == 0 {
				retMap[lower] = higher
			}
		}
	}
	return
}

func checkEmirp(primes []int64, p int64) (lower int64, higher int64, isEmirp bool) {
	s := fmt.Sprintf("%v", p)
	newS := commonstringutils.ReverseString(s)
	if s != newS {
		newP, _ := strconv.ParseInt(newS, 10, 64)
		if contains(primes, newP) {
			if p < newP {
				return p, newP, true
			}
			return newP, p, true
		}
	}
	return 0, 0, false
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findPalendromes(primes []int64) (retArray []int64) {
	for _, p := range primes {
		s := fmt.Sprintf("%v", p)
		if isPalendrome(s) {
			retArray = append(retArray, p)
		}
	}
	return
}

func isPalendrome(s string) bool {
	if s == commonstringutils.ReverseString(s) {
		return true
	}
	return false
}

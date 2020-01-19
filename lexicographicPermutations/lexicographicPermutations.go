package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation
of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically,
we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

func main() {
	getNthLexiPermutation([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, int64(1000000))
}

func getNthLexiPermutation(charSet []string, n int64) (mutation []string, count int64) {
	fmt.Println(fmt.Sprintf("Finding the %v lexicological permutation of %s", n, charSet))

	mutation = make([]string, len(charSet))
	copy(mutation, charSet)
	sort.Strings(mutation)
	endMutation := make([]string, len(mutation))
	copy(endMutation, mutation)
	sort.Sort(sort.Reverse(sort.StringSlice(endMutation)))

	for count = int64(1); count < n && !stringSliceExactMatch(mutation, endMutation); count++ {
		mutation = getNextMutation(mutation)
	}

	if count < n {
		fmt.Println(fmt.Sprintf("Unable to continue at %v lexicological permutation %s", count, mutation))
	} else {
		fmt.Println(fmt.Sprintf("The %v lexicological permutation is %s", count, mutation))
	}

	return mutation, count
}

func stringSliceExactMatch(a []string, b []string) bool {
	as := strings.Join(a, "")
	bs := strings.Join(b, "")
	return as == bs
}

func getNextMutation(oldMutation []string) (newMutation []string) {
	_, indexF := findLastSmallestCharacter(oldMutation)
	_, indexS := findNextSmallestRightCharacter(oldMutation, indexF)
	interumSlice := swapCharacters(oldMutation, indexF, indexS)
	newMutation = sortAfter(interumSlice, indexF)
	return
}

func findLastSmallestCharacter(ss []string) (char string, index int) {
	if len(ss) == 0 {
		return "", 0
	}

	for index := len(ss) - 1; index >= 0; index-- {
		if index > 0 && ss[index] > ss[index-1] {
			return ss[index-1], index - 1
		}
	}
	return ss[0], 0
}

func findNextSmallestRightCharacter(ss []string, indexF int) (char string, indexS int) {

	if len(ss) == 0 {
		return "", 0
	}

	var previousSmallest string = ""
	var previousIndex int = 0

	if len(ss)-1 == indexF {
		previousSmallest, previousIndex = ss[indexF], indexF
	} else {
		previousSmallest, previousIndex = ss[indexF+1], indexF+1
	}

	for indexS := indexF + 1; indexS <= len(ss)-1; indexS++ {

		if ss[indexS] < previousSmallest && ss[indexS] > ss[indexF] {
			previousSmallest, previousIndex = ss[indexS], indexS
		}
	}
	return previousSmallest, previousIndex
}

func swapCharacters(ss []string, indexF int, indexS int) (newSlice []string) {
	if len(ss) == 0 {
		return ss
	}

	newSlice = make([]string, len(ss))
	copy(newSlice, ss)
	newSlice[indexF] = ss[indexS]
	newSlice[indexS] = ss[indexF]
	return
}

func sortAfter(ss []string, index int) []string {
	if len(ss) == 0 {
		return ss
	}

	front := ss[:index+1]
	back := ss[index+1:]
	sort.Strings(back)

	return append(front, back...)
}

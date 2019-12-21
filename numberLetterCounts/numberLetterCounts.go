package main

/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there
are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many
letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23
letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out
numbers is in compliance with British usage.
*/

import (
	"fmt"
	"projectEuler/common/num2words"
	"strings"
)

func main() {
	//getNumberLetterCounts(5) //provided example
	getNumberLetterCounts(1000)
}

func getNumberLetterCounts(max int) (words string, count int) {
	fmt.Println(fmt.Sprintf("calculating the words for numbers up to: %v", max))

	words, count = processNums(max)
	fmt.Println(fmt.Sprintf("the words are: %s", words))
	fmt.Println(fmt.Sprintf("the count of letters is: %v", count))
	return
}

func processNums(max int) (words string, count int) {
	for num := 1; num <= max; num++ {
		words = words + " " + num2words.Convert(num)
		words = strings.TrimSpace(words)
	}
	packedWords := strings.Replace(words, " ", "", -1)
	count = len(packedWords)
	return
}

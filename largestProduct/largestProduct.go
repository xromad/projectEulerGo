package main

import (
	"fmt"
	"projectEuler/common/commonintutils"
	"strings"
)

/*
The four adjacent digits in the 1000-digit number that have the greatest product
are 9 × 9 × 8 × 9 = 5832.

73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450

Find the thirteen adjacent digits in the 1000-digit number that have the greatest product.
What is the value of this product?
*/

func main() {
	//var length int = 4 //5832 the provided example
	var length int = 13 //23514624000 the provided question

	sequence := getStringLiteral()
	int64Slice := stringToInt64Slice(sequence)

	fmt.Println(fmt.Sprintf("Finding largest product of %v digits in the sequence", length))
	digitsSlice, product := findLargestProduct(length, int64Slice)
	fmt.Println(fmt.Sprintf("Largest product of %v sequential digits %v is %v",
		length,
		digitsSlice,
		product,
	))
}

func findLargestProduct(length int, int64Slice []int64) (factorSlice []int64, product int64) {
	for i := range int64Slice {
		subSlice := getNextSlice(i, length, int64Slice)
		subProduct := int64(1)

		for _, n := range subSlice {
			subProduct = subProduct * n
		}

		if subProduct > product {
			product = subProduct
			factorSlice = subSlice
		}
	}
	return
}

func getNextSlice(start int, length int, int64Slice []int64) (subSlice []int64) {
	end := start + length
	if end > len(int64Slice) {
		return int64Slice[start:len(int64Slice)]
	}
	return int64Slice[start:end]
}

func getStringLiteral() string {
	return strings.ReplaceAll(
		`73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`,
		"\n", "")
}

func stringToInt64Slice(s string) (int64Slice []int64) {
	for _, s := range s {
		int64Slice = append(int64Slice, commonintutils.AtoInt64(string(s)))
	}
	return int64Slice

}

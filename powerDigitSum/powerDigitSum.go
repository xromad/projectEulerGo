package main

/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
What is the sum of the digits of the number 2^1000?
*/

import (
	"fmt"
	"math"
	"projectEuler/common/commonintutils"
	"strconv"
)

func main() {
	//finalSum := powerDigitSum(float64(15)) //provided example
	finalSum := powerDigitSum(float64(1000))
	fmt.Println(fmt.Sprintf("sum: %v", finalSum))
}

func powerDigitSum(pow float64) int {
	fmt.Println(fmt.Sprintf("finding sum of the digits of : 2^%v", pow))

	result := math.Pow(2, pow)
	fmt.Println(fmt.Sprintf("2^%.f = %.f", pow, result))

	digits := getDigits(result)
	fmt.Println(fmt.Sprintf("digits: %v", digits))

	return commonintutils.SumIntSlice(digits)
}

func getDigits(num float64) (result []int) {
	numString := commonintutils.F64toa(num)
	for _, r := range numString {
		digit, _ := strconv.Atoi(string(r))
		result = append(result, digit)
	}
	return
}

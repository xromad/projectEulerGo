package main

/*
 * given numbers between 1 and 100
 * produce an equation in the form of a+b...=c
 * where the additive numbers are divisible by 3 or 5 and the sum is at the end
 * example for the numbers 1 through 10 the solution would be 3+5+6+9=23
 */

import (
	"fmt"
	"projectEuler/common/commonintutils"
	"strconv"
	"strings"
)

func main() {
	start := 1
	end := 100

	fmt.Println(fmt.Sprintf("Running equation generator for numbers between: %v and %v", start, end))
	fmt.Println(genEquation(start, end))
}

func genEquation(begin, end int) string {
	if begin < 1 {
		return "Start must be >= 1"
	}
	if end > 100 {
		return "End must be <= 100"
	}

	var values []int
	for i := begin; i < end; i++ {
		if eqDecider(i) {
			values = append(values, i)
		}
	}

	return buildString(values)
}

func buildString(nums []int) string {
	if len(nums) == 0 {
		return "0=0"
	}

	sVals := commonintutils.IntSliceToStringSlice(nums)

	builder := strings.Builder{}
	builder.WriteString(strings.Join(sVals, "+"))
	builder.WriteString("=")
	builder.WriteString(strconv.Itoa(commonintutils.SumIntSlice(nums)))

	return builder.String()
}

func eqDecider(i int) bool {
	return i%3 == 0 || i%5 == 0
}

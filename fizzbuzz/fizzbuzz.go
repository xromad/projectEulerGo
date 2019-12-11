package main

/*
 * FizzBuzz problem
 * for all integers between 1 and 100
 * if i divisible by 3 then print "fizz"
 * if i divisible by 5 then print "buzz"
 * if i divisible by 3 and 5 then print "fizzbuzz"
 * else print i
 */

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("running: 1 -> 100")
	fmt.Println(genOutput(1, 100))
}

func genOutput(begin, end int) string {
	builder := strings.Builder{}
	for i := begin; i <= end; i++ {
		if i == begin {
			builder.WriteString(getFizzBuzz(i))
		} else {
			builder.WriteString(" " + getFizzBuzz(i))
		}
	}
	return builder.String()
}

func getFizzBuzz(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "fizzbuzz"
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	default:
		return strconv.Itoa(i)
	}
}

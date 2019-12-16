// Package num2words - a utility package to convert numbers to words
package num2words

//NOTE: this only does up to thousands at this time
//NOTE: this only does positives at this time
//TODO: use recursion and a builder to make this much better

import "math"

// Convert - convert an int into a set of words
func Convert(n int) string {
	switch {
	case n > 0 && n < 10:
		return ones(n)
	case n >= 10 && n < 100:
		return tens(n)
	case n >= 100 && n < 1000:
		return hundreds(n)
	case n >= 1000 && n < 10000:
		return thousands(n)
	default:
		return ""
	}
}

//given thousands convert to string
func thousands(n int) string {
	if n >= 1000 && n < 10000 {
		hundredsComponent := n % 1000
		thousandsComponent := int(math.Floor(float64(n) / 1000))
		if hundredsComponent == 0 {
			return ones(thousandsComponent) + " thousand"
		}
		return ones(thousandsComponent) + " thousand " + hundreds(hundredsComponent)
	}
	return hundreds(n)
}

//given hundreds convert to string
func hundreds(n int) string {
	if n >= 100 && n < 1000 {
		tensComponent := n % 100
		hundredsComponent := int(math.Floor(float64(n) / 100))
		if tensComponent == 0 {
			return ones(hundredsComponent) + " hundred"
		}
		return ones(hundredsComponent) + " hundred " + tens(tensComponent)
	}
	return tens(n)
}

//given tens convert to string
func tens(n int) string {
	conversionMap := getMap()
	switch {
	case n >= 10 && n < 20:
		return (*conversionMap)[n]
	case n >= 20 && n < 100:
		onesComponent := n % 10
		tensComponent := n - onesComponent
		if onesComponent == 0 {
			return (*conversionMap)[tensComponent]
		}
		return (*conversionMap)[tensComponent] + " " + ones(onesComponent)
	default:
		return ones(n)
	}
}

//given a single digit convert it to string
func ones(n int) string {
	if n < 10 {
		conversionMap := getMap()
		return (*conversionMap)[n]
	}
	return ""
}

func getMap() *map[int]string {
	return &map[int]string{
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
		20: "twenty", 30: "thirty", 40: "fourty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
	}

}

// Package commonintutils - common utilities for working with ints
package commonintutils

import (
	"fmt"
	"strconv"
)

/*
 * int utils
 */

// ProductIntSlice - returns the product of the values in the slice
func ProductIntSlice(factors []int) (product int) {
	if len(factors) == 0 {
		return 0
	}

	product = 1
	for _, i := range factors {
		product = product * i
	}
	return
}

// IntSliceToStringSlice  - converts a slice of ints to a slice of strings
func IntSliceToStringSlice(nums []int) (sVals []string) {
	for _, num := range nums {
		sVals = append(sVals, strconv.Itoa(num))
	}
	return
}

// SumIntSlice - adds all values in the slice and returns the sum
func SumIntSlice(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

// IntSliceEquals - tests to see if two int Slices are equal
func IntSliceEquals(aSlice []int, bSlice []int) bool {
	if len(aSlice) != len(bSlice) {
		return false
	}
	for _, v := range aSlice {
		if !IntSliceContains(bSlice, v) {
			return false
		}
	}
	return true
}

// IntSliceContains - tests to see if the int slice contains the value
func IntSliceContains(bSlice []int, v int) bool {
	for _, c := range bSlice {
		if c == v {
			return true
		}
	}
	return false
}

/*
 * uint Utils
 */

// UIntSliceEquals - tests to see if two uint Slices are equal
func UIntSliceEquals(aSlice []uint, bSlice []uint) bool {
	if len(aSlice) != len(bSlice) {
		return false
	}
	for _, v := range aSlice {
		if !UIntSliceContains(bSlice, v) {
			return false
		}
	}
	return true
}

// UIntSliceContains - tests to see if the int slice contains the value
func UIntSliceContains(bSlice []uint, v uint) bool {
	for _, c := range bSlice {
		if c == v {
			return true
		}
	}
	return false
}

/*
 * float64 Utils
 */

// F64toa - convert float64 number to string
func F64toa(num float64) string {
	return fmt.Sprintf("%.f", num)
}

/*
 * int64 utils
 */

// I64toA - convert int64 number to string
func I64toA(num int64) string {
	return fmt.Sprintf("%v", num)
}

// AtoInt64 - convert string number to int64
func AtoInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// Int64SliceEquals - tests to see if two int64 Slices are equal
func Int64SliceEquals(aSlice []int64, bSlice []int64) bool {
	if len(aSlice) != len(bSlice) {
		return false
	}
	for _, v := range aSlice {
		if !Int64SliceContains(bSlice, v) {
			return false
		}
	}
	return true
}

// Int64SliceContains - tests to see if the int64 slice contains the value
func Int64SliceContains(bSlice []int64, v int64) bool {
	for _, c := range bSlice {
		if c == v {
			return true
		}
	}
	return false
}

// SumInt64Slice - adds all values in the slice and returns the sum
func SumInt64Slice(nums []int64) (sum int64) {
	for _, t := range nums {
		sum += t
	}
	return
}

// Int64SliceToStringSlice  - converts a slice of ints to a slice of strings
func Int64SliceToStringSlice(nums []int64) (sVals []string) {
	for _, num := range nums {
		sVals = append(sVals, fmt.Sprintf("%v", num))
	}
	return
}

// ProductInt64Slice - returns the product of the values in the slice
func ProductInt64Slice(factors []int64) (product int64) {
	if len(factors) == 0 {
		return 0
	}

	product = 1
	for _, i := range factors {
		product = product * i
	}
	return
}

// ProductInt64Range - returns the product of the values in the range
func ProductInt64Range(low int64, high int64) (product int64) {
	if low < 1 || high < 1 || high < low {
		return 0
	}
	product = 1
	for i := low; i <= high; i++ {
		product = product * i
	}
	return
}

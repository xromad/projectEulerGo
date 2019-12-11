// Package commonstringutils - a reversing utilty package
package commonstringutils

// ReverseString - uses runes to reverse a unicode string
func ReverseString(s string) string {
	runes := []rune(s)
	var rOut []rune

	for _, r := range runes {
		//prepend is append with your output last...
		//but wierder because you have to convert your element into a slice first
		rOut = append([]rune{r}, rOut...)
	}

	return string(rOut)
}

// StringSliceCompare - tests to see if the two string slices are the same
func StringSliceCompare(s1 []string, s2 []string) bool {
	if StringHasVals(s1, s2) {
		return true
	}
	return false
}

// StringHasVals - tests to see if the two string slices have the same values
func StringHasVals(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, e := range s1 {
		if !StringSliceContains(s2, e) {
			return false
		}
	}
	return true
}

// StringSliceContains - tests to see if the string slice contains the string
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//IsPalendrome - tests to see if a string is a palendrome
func IsPalendrome(s string) bool {
	if s == ReverseString(s) {
		return true
	}
	return false
}

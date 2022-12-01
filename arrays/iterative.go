package arrays

import lib "github.com/tlaceby/go-utils"

// Get the accumulated value from an array of numeric values
func Accumulate[T lib.Number](array []T) T {
	var sum T

	for _, val := range array {
		sum += val
	}

	return sum
}

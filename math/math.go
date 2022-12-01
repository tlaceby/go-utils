package math

import (
	lib "github.com/tlaceby/go-utils"
)

// Returns the largest value inside an array of numeric values
func Max[T lib.Number](array []T) T {
	var largest T
	for _, val := range array {
		if val > largest {
			largest = val
		}
	}
	return largest
}

// Returns the smallest value inside an array of numeric values
func Min[T lib.Number](array []T) T {
	var smallest T
	for _, val := range array {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}

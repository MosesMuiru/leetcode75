package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	fmt.Println(first)
	for _, num := range nums {
		if num <= first {
			first = num // smallest so far
		} else if num <= second {
			second = num // second smallest
		} else {
			// Found a number greater than both first and second
			return true
		}
	}

	return false
}

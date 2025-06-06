package main

import (
	"slices"
	"testing"
)

func TestGreatestCandies(t *testing.T) {
	candies, extracandies := []int{4, 2, 1, 1, 2}, 1
	output := []bool{true, false, false, false, false}

	results := GreatestCand(candies, extracandies)
	if !slices.Equal(results, output) {
		t.Errorf("Budaa incorrect, got, %v, expected --> %v", results, output)

	}

}

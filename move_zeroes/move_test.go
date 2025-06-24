package main

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	input := []int{0, 1, 0, 3, 12}
	output := []int{1, 3, 12, 0, 0}

	results := MoveZeroes(input)

	if !slices.Equal(results, output) {
		t.Errorf("Budaa got %v, was expecting %v", output, results)
	}

}

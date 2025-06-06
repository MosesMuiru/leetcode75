package main

import (
	"slices"
	"testing"
)

func TestInter(t *testing.T) {

	one := []int{1, 2, 3}
	two := []int{2, 4, 6}
	expected := []int{2, 2}

	results := Inter(one, two)
	if slices.Equal(one, two) {
		t.Errorf("got the following %v,  expecting %v", results, expected)
	}
}

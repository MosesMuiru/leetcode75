package main

import (
	"testing"
)

func TestCompressString(t *testing.T) {
	input := []string{"a", "a", "b", "b", "c", "c"}
	//output := []string{"a", "2", "b", "2", "c", "3"}
	re := 6

	results := CompressString(input)

	if results != re {
		t.Errorf("Buda was expecting %v, got %v", re, results)
	}
}

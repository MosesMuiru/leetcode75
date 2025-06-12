package main

import "testing"

func TestReverseWord(t *testing.T) {
	input := "a good   example"
	output := "example good a"

	results := ReverseWord(input)

	if results != output {
		t.Errorf("Buda got this -- %v, was expecting thi--%v", results, output)
	}
}

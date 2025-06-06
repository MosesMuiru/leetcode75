package main

import "testing"

func TestMergeStrings(t *testing.T) {
	word1, word2 := "abc", "pqrs"

	results := MergeStrings(word1, word2)
	should := "apbqcrs"

	if results != should {
		t.Errorf("Budaa incorrect, got, %s, expected --> %s", results, should)
	}
}

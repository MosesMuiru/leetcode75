package main

import "testing"

func TestSubsequence(t *testing.T) {
	s := "aaaaaa"
	v := "bbaaaa"

	results := Subsequence(s, v)

	if results != true {
		t.Errorf("Buda was expecing %v, and go this %v", s, results)
	}

}

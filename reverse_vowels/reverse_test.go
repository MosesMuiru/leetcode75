package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	Input := "IceCreAm"

	Output := "AceCreIm"
	results := Reverse(Input)
	fmt.Println("resultsss_>", results)

	if results != Output {
		t.Errorf("Budaa i was exepcting %s, and i have got %s", Output, results)
	}

}

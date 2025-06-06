package main

import (
	"fmt"
	"testing"
)

func TestCommonDivisor(t *testing.T) {
	str1, str2 := "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	expects := "TAUXX"
	fmt.Println(str2)
	results := CommonDivisor(str1, str2)

	if results != expects {
		t.Errorf("Buda errror, got %v, and was expecting -->%v", results, expects)
	}

}

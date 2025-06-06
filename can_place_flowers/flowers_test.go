package main

import (
	"fmt"
	"testing"
)

type testData struct {
	flowerbed []int
	n         int
	response  bool
}

func TestCanPlaceFlowers(t *testing.T) {
	data := []testData{
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 1, response: true},
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 2, response: false},
		{flowerbed: []int{1, 0, 1, 0, 1, 0, 1}, n: 1, response: false},
		{flowerbed: []int{1, 0, 1, 0, 1, 0, 1}, n: 0, response: true},
		{flowerbed: []int{1, 0, 0, 0, 1, 0, 1}, n: 1, response: true},
		{flowerbed: []int{1, 0, 0, 0, 0, 1}, n: 2, response: false},
		{flowerbed: []int{1, 0, 0, 1, 0, 0}, n: 2, response: false},
	}

	for _, p := range data {
		output := PlaceFlowers(p.flowerbed, p.n)

		if output != p.response {
			fmt.Println("THis is test data", p.flowerbed)
			t.Errorf("Buda radaaa, i got %v, and was expecting %v, n: %v, from this %v", output, p.response, p.n, p.flowerbed)
		}

	}
}

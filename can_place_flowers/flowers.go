package main

import "fmt"

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// calculate the number of empty spaces in the flowerbed
	empty := n*2 + 1

	output := false
	if n == 0 {
		output = true
	}
	spaces := 0
	flowers := 0

	// check if there are empty beds

	for i := 0; i < len(flowerbed); i++ {
		fmt.Println("ii", i)
		if flowerbed[i] == 0 {
			spaces += 1
		}

		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {

			if n != 0 {
				n -= 1
				flowers += 1
			}

			fmt.Printf("left side %v, middle %v, right side %v", flowerbed[i], flowerbed[i-1], flowerbed[i+1])
			output = true

		}
	}

	fmt.Println("This is what is left for n", n)

	if spaces != empty {
		output = false
	}
	if flowers == n {
		output = true
	}

	fmt.Println(empty)
	fmt.Println("spacess", spaces)

	return output
}

func PlaceFlowers(flowerBeds []int, n int) bool {
	output := false

	f_planted := 0
	for i := 0; i < len(flowerBeds); i++ {
		// plant the flower

		fmt.Println("i -.", i)
		if flowerBeds[i] == 0 &&
			(i == 0 || flowerBeds[i-1] == 0) &&
			(i == len(flowerBeds)-1 || flowerBeds[i+1] == 0) {

			// you can now plant a flower
			flowerBeds[i] = 1
			f_planted += 1
		}

	}

	fmt.Println("Planted flower", f_planted)
	fmt.Println(" flower, reamining", n)

	if f_planted >= n {
		output = true
	}

	return output
}

func main() {
	flowerbeds := []int{0, 0, 1, 0, 0}
	response := PlaceFlowers(flowerbeds, 1)
	fmt.Println("respnse", response)

}

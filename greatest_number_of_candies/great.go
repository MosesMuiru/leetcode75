package main

// n kids with candies
// extracandies

// return a bolean arra results of the length with

func GreatestCand(candies []int, extraCandies int) []bool {
	max := candies[0]

	// find the max number of candies any kid currently has
	for _, value := range candies {
		if value > max {
			max = value
		}
	}

	// check each kid after dispersoal

	output := make([]bool, len(candies))
	for i, value := range candies {
		if value+extraCandies >= max {
			output[i] = true
		} else {
			output[i] = false
		}
	}

	return output
}

func main() {

	candies, extracandies := []int{2, 8, 7}, 1
	_ = GreatestCand(candies, extracandies)
}

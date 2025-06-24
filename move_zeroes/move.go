package main

func MoveZeroes(nums []int) []int {

	count := 0
	zeroes := 0

	for _, value := range nums {
		if value != 0 {
			nums[count] = value
			count++
			continue
		}
		zeroes++
	}
	if len(nums) != 1 && len(nums) != 0 && zeroes != 0 {

		for i := len(nums) - 1; i >= 0; i-- {
			nums[i] = 0
			zeroes--
			if zeroes == 0 {
				break
			}
		}
	}

	return nums

}

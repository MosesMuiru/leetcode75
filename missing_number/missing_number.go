package missingnumber

func missingNumber(nums []int) int {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1]-nums[i] != 1 {
			return nums[i] + 1
		} else if len(nums) == 2 {
			return 2
		}
	}
	return nums[len(nums)-1]
}

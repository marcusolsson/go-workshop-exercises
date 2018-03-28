package slices

func maximum(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// BONUS!
func rangeMaximum(nums []int) int {
	var max int
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max
}

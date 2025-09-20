package main

func max(nums ...int) int {
	var high int
	if len(nums) > 0 {
		high = nums[0]
	}
	for _, num := range nums {
		if num > high {
			high = num
		}
	}
	return high
}

func min(nums ...int) int {
	var low int
	if len(nums) > 0 {
		low = nums[0]
	}
	for _, num := range nums {
		if num > low {
			low = num
		}
	}
	return low
}

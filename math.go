package goutils

import "math"

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}

	return max
}

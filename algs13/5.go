package algs13

import "sort"

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sum := 0
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 1; i < len(nums); i++ {
		if nums[i-1]-nums[i] > sum {
			sum = nums[i-1] - nums[i]
		}
	}
	return sum
}

package algs13

import (
	"sort"
)

func FindKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

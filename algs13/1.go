package algs13

import "sort"

/*
* 16. 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
*
 */

func ThreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	n, bestnum := len(nums), nums[0]+nums[1]+nums[2] // 初始化为前三元素的值，避免任何个

	for i := 0; i < n-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(bestnum-target) {
				bestnum = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return bestnum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

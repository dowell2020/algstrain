/*
 * @Author: dowell87
 * @Date: 2021-07-22 21:26:47
 * @Descripttion:
 * @LastEditTime: 2021-07-22 22:00:42
 */
package algs1

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func ThreeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	var data = 0
	n := len(nums)
	// data = abs(data)
	// 根据差值的绝对值来更新答案
	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		if n1 > target {
			break
		}
		// if i > 0 && n1 == nums[i-1] {
		// 	continue
		// }
		l, r := i+1, len(nums)-1
		for l < r || data == target {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) <
				math.Abs(float64(data-target)) { // 判断谁离target更近
				data = sum
			}
			// fmt.Println(data)

			l++
			r--
		}
	}
	return data
}

// @lc code=end

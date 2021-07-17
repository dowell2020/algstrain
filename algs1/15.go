/*
 * @Author: dowell87
 * @Date: 2021-07-15 23:03:33
 * @Descripttion:
 * @LastEditTime: 2021-07-15 23:33:11
 */
package algs1

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	data := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// for j := i + 1; i < len(nums)-1; i++ {

		// }
	}

	return data
}

// @lc code=end

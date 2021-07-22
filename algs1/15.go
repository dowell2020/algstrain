/*
 * @Author: dowell87
 * @Date: 2021-07-15 23:03:33
 * @Descripttion:
 * @LastEditTime: 2021-07-22 21:25:24
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
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				data = append(data, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return data
}

// @lc code=end

//
// 写的质量太差了，重新在写一次
// func ThreeSum(nums []int) [][]int {
// 	n := len(nums)
// 	if n < 3 {
// 		return nil
// 	}
// 	sort.Ints(nums)
// 	data := make([][]int, 0)
// 	z := 1
// 	// 全部循环根据角标判断是否重复
// 	for i := 0; i < len(nums)-2; i++ {
// 		if nums[i] > 0 {
// 			break
// 		}
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}
// 		for n := i + 1; n < len(nums); n++ {
// 			for j := len(nums) - 1; n < j; j-- {
// 				if nums[i]+nums[n]+nums[j] == 0 {
// 					// 产生的三个数的值是否一致
// 					// NumI := [...]int{nums[i], nums[i], nums[i]}
// 					for k := 0; k < len(data); k++ {
// 						if nums[i] == data[k][0] && nums[n] == data[k][1] {
// 							z = 0
// 						}
// 					}
// 					if z == 1 {
// 						data = append(data, []int{nums[i], nums[n], nums[j]})
// 					}
// 					z = 1
// 					// 	// fmt.Println(111111 + i)
// 				}

// 			}

// 		}
// 	}

// 	return data
// }

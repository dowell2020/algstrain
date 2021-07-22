/*
 * @Author: dowell87
 * @Date: 2021-07-08 10:21:34
 * @Descripttion:
 * @LastEditTime: 2021-07-21 20:54:34
 */
package algs1

import "fmt"

func TttTest() {
	fmt.Println("aa")
}

// 大部分是抄袭的
// 后面再做一次

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func LengthOfLongestSubstring(s string) int {
	cot := len(s)
	if cot == 0 {
		return 0
	}
	tempMap := make(map[uint8]int)
	l, r, maxLen := 0, 1, 1
	tempMap[s[l]] = l
	for l < r && r < cot {
		val, ok := tempMap[s[r]]
		if ok && val >= l {
			l = val + 1
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}

		tempMap[s[r]] = r
		r++
	}

	return maxLen
}

// @lc code=end

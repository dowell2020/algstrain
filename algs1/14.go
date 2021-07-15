/*
 * @Author: dowell87
 * @Date: 2021-07-15 21:00:42
 * @Descripttion:
 * @LastEditTime: 2021-07-15 23:02:49
 */

package algs1

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	pre := strs[0]
	for _, k := range strs {
		for strings.Index(k, pre) != 0 {
			if len(pre) == 0 {
				return ""
			}
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}

// @lc code=end

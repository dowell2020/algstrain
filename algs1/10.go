/*
 * @Author: dowell87
 * @Date: 2021-07-15 13:36:44
 * @Descripttion:
 * @LastEditTime: 2021-07-15 16:24:43
 */
package algs1

/**
 * @description: 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 * @param {*}
 * @return {*}
*/

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func IsMatch(s string, p string) bool {
	if len(s) > len(p) {
		return false
	}
	// p := /p/
	// res, err := regexp.MatchString(s, p)
	// if err != nil {
	// 	return false
	// }
	return true
}

// @lc code=end

// 题意还没有理解清楚
// 放一放再做

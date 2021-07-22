/*
 * @Author: dowell87
 * @Date: 2021-07-08 16:27:40
 * @Descripttion:
 * @LastEditTime: 2021-07-21 20:54:41
 */
/*
 * @Author: dowell87
 * @Date: 2021-07-08 16:26:25
 * @Descripttion:
 * @LastEditTime: 2021-07-08 16:27:22
 */
package algs1

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var stringlen int = len(s)
	if stringlen <= 1 {
		return s
	}
	var i int = 0
	var maxLen = 0
	var res = ""
	for i < stringlen {
		var pointerL = i
		var pointerR = i
		for pointerL >= 0 && s[pointerL] == s[i] {
			pointerL--
		}
		for pointerR < stringlen && s[pointerR] == s[i] {
			pointerR++
		}
		var nextPoint = pointerR
		for pointerL >= 0 && pointerR < stringlen && s[pointerR] == s[pointerL] {
			pointerL--
			pointerR++
		}
		var tmpMaxLen = pointerR - pointerL - 1
		if tmpMaxLen > maxLen {
			res = s[pointerL+1 : pointerR]
			maxLen = tmpMaxLen
		}
		i = nextPoint
	}
	return res
}

// @lc code=end

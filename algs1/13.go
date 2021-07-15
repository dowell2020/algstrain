/*
 * @Author: dowell87
 * @Date: 2021-07-15 18:21:19
 * @Descripttion:
 * @LastEditTime: 2021-07-15 20:58:57
 */
package algs1

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func RomanToInt(s string) int {
	// var romanS = []string{"I", "V", "X", "L", "C", "D", "M"}
	// var romanI = []int{1, 5, 10, 50, 100, 500, 1000}
	// var num = 0
	// // var TempInt int
	// // 一直减
	// j := 0
	// i := 7
	// TempNum := 0
	// for len(s) > j {
	// 	i = len(romanI) - 1
	// 	for i >= 0 {
	// 		if string(s[j]) == romanS[i] {
	// 			if romanI[i] > romanI[TempNum] {
	// 				num = num + romanI[i]
	// 			} else {
	// 				num = num - romanI[i]
	// 			}
	// 			TempNum = i
	// 			j++
	// 			break
	// 		}

	// 		i--
	// 	}
	// }
	// return num
	var mp = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var num int

	tmp := 1000
	for i := 0; i < len(s); i++ {
		now := mp[s[i]]
		if tmp < now {
			num -= 2 * tmp
		}
		num += now
		tmp = now
	}
	return num
}

// @lc code=end

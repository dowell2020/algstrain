/*
 * @Author: dowell87
 * @Date: 2021-07-08 22:44:16
 * @Descripttion:
 * @LastEditTime: 2021-07-09 01:00:27
 */
package leetcode

import (
	"math"
	"strings"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func MyAtoi(s string) int {
	str := strings.TrimSpace(s)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result

	// cst := len(s)
	// if cst < 1 {
	// 	return 0
	// }
	// i := 0
	// z := 0
	// j := 0
	// num := 0
	// for i < len(s) {
	// 	fmt.Println(s[i])
	// 	if s[i] == 32 {
	// 		j = -1
	// 	}
	// 	if s[i] == 45 {
	// 		z = -1
	// 		if j == 0 || j == -1 {
	// 			j++
	// 		}
	// 	}
	// 	if s[i] == 43 {
	// 		if j == 0 || j == -1 {
	// 			j++
	// 		}
	// 	}
	// 	if s[i] < 58 && s[i] > 47 {
	// 		t, _ := strconv.Atoi(string(s[i]))
	// 		if z < 0 {
	// 			num = num*10 - t
	// 		} else {
	// 			num = num*10 + t
	// 		}
	// 		j++
	// 	}
	// 	if j == i || num > 0 {
	// 		i = len(s)
	// 		fmt.Println("aa")
	// 	}
	// 	i++
	// }
	// if num > 2147483647 {
	// 	num = 2147483647
	// }
	// if num < -2147483648 {
	// 	num = -2147483648
	// }
	// return num
}

// @lc code=end

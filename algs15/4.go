/*
 * @Author: dowell87
 * @Date: 2021-11-18 15:29:57
 * @Descripttion:
 * @LastEditTime: 2021-11-18 17:09:57
 */
package algs15

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */

// @lc code=start
func ComplexNumberMultiply(num1 string, num2 string) string {
	var x, y, m, n int
	fmt.Sscanf(num1, "%d+%di", &x, &y)
	fmt.Sscanf(num2, "%d+%di", &m, &n)
	return fmt.Sprintf("%d+%di", (x*m - y*n), (m*y + x*n))
}

// @lc code=end

/* 需要了解一些原生的方法组件 */

/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
func CarPooling(trips [][]int, capacity int) bool {
	stops := make([]int, 1001)
	for _, t := range trips {
		stops[t[1]] += t[0]
		stops[t[2]] -= t[0]
	}
	for i := 0; capacity >= 0 && i < 1001; i++ {
		capacity -= stops[i]
	}
	return capacity >= 0
}

// @lc code=end

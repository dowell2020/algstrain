/*
 * @Author: dowell87
 * @Date: 2021-11-11 11:37:47
 * @Descripttion:
 * @LastEditTime: 2021-11-11 17:33:58
 */
package algs15

import (
	"strconv"
)

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	str := ""
	new1, new2 := len(num1), len(num2)
	ansArr := make([]int, new1+new2)
	for i := new1 - 1; i >= 0; i-- {
		for j := new2 - 1; j >= 0; j-- {
			ansArr[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
			ansArr[i+j] += ansArr[i+j+1] / 10
			ansArr[i+j+1] %= 10
		}
	}
	for k, v := range ansArr {
		if k == 0 {
			if v != 0 {
				str = str + strconv.Itoa(v)
			}
		} else {
			str = str + strconv.Itoa(v)
		}
	}
	return str
}

func FindDiagonalOrder2(mat [][]int) []int {
	if mat == nil || len(mat) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])

	result := make([]int, m*n)
	row, col, d := 0, 0, 0
	dirs := [][]int{{-1, 1}, {1, -1}}

	for i := 0; i < m*n; i++ {
		result[i] = mat[row][col]
		row += dirs[d][0]
		col += dirs[d][1]

		if row >= m {
			row = m - 1
			col += 2
			d = 1 - d
		}
		if col >= n {
			col = n - 1
			row += 2
			d = 1 - d
		}
		if row < 0 {
			row = 0
			d = 1 - d
		}
		if col < 0 {
			col = 0
			d = 1 - d
		}
	}

	return result
}

func FindDiagonalOrder(mat [][]int) []int {
	if mat == nil || len(mat) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])
	r, c := 0, 0
	arr := make([]int, m*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = mat[r][c]
		if (r+c)%2 == 0 { // moving up
			if c == n-1 {
				r++
			} else if r == 0 {
				c++
			} else {
				r--
				c++
			}
		} else { // moving down
			if r == m-1 {
				c++
			} else if c == 0 {
				r++
			} else {
				r++
				c--
			}
		}
	}
	return arr
}

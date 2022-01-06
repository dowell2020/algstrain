/*
 * @Author: dowell87
 * @Date: 2021-11-03 11:06:58
 * @Descripttion:
 * @LastEditTime: 2021-11-03 17:05:35
 */
package algs14

/**
 * @description:48. 旋转图像
 * @param {[][]int} matrix
 * @return {*}
 */

func Rotate(matrix [][]int) {
	quadrats := len(matrix) / 2
	for q := 0; q < quadrats; q++ {
		beg := 0 + q
		end := len(matrix) - 1 - q
		swaps := len(matrix) - 1 - q*2
		for i := 0; i < swaps; i++ {
			topLeft := matrix[beg][beg+i]
			matrix[beg][beg+i] = matrix[end-i][beg]
			matrix[end-i][beg] = matrix[end][end-i]
			matrix[end][end-i] = matrix[beg+i][end]
			matrix[beg+i][end] = topLeft
		}
	}
}

/**
 * @description: 36. 有效的数独
 * @param {[][]byte} board
 * @return {*}
 */

func IsValidSudoku(board [][]byte) bool {
	// seen := make(n,)
	mp1 := map[byte]bool{}
	mp2 := map[byte]bool{}
	mp3 := map[byte]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 记录一行数据
			// 记录一列数据
			// 记录块状数据
			if board[i][j] != '.' {
				if mp1[board[i][j]] || mp2[board[j][i]] || mp3[board[(i%3)*3+j%3][(i/3)*3+j/3]] {
					return false
				}

			}
		}
	}
	return true
}

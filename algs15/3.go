/*
 * @Author: dowell87
 * @Date: 2021-11-18 14:27:18
 * @Descripttion:
 * @LastEditTime: 2021-11-18 15:28:23
 */
package algs15

func GameOfLife(board [][]int) {
	if board == nil || len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 循环每一个数字
			// 统计细胞周边的细胞
			count := liveNeighbors(board, m, n, i, j)
			if board[i][j] == 1 && count >= 2 && count <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && count == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

// 获取八个角度的方法可以记录下来
func liveNeighbors(board [][]int, m, n, i, j int) int {
	lives := 0
	for x := max(i-1, 0); x <= min(i+1, m-1); x++ {
		for y := max(j-1, 0); y <= min(j+1, n-1); y++ {
			lives += board[x][y] & 1
		}
	}
	lives -= board[i][j] & 1
	return lives
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

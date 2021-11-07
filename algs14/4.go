/*
 * @Author: dowell87
 * @Date: 2021-11-07 19:38:00
 * @Descripttion:
 * @LastEditTime: 2021-11-07 22:23:15
 */
package algs14

func tictactoe(board []string) (res string) {
	// 获取棋盘的大小
	num := len(board)
	// 判断是否还有下棋位置
	isFull := true
	// 统计落子的数量
	rowS := make([]string, num)
	colS := make([]string, num)
	// 对角线A
	rowA := board[0][:1]
	// 对角线B
	rowB := board[0][num-1 : num]
	for i := 0; i < num; i++ {
		rowS[i] = board[i][:1]
		for j := 0; j < num; j++ {
			// 行
			if j < num-1 {
				if board[i][j:j+1] != board[i][j+1:j+2] {
					rowS[i] = ""
				}
			}
			// 列
			if i == 0 {
				colS[j] = board[i][j : j+1]
			}
			if i > 0 {
				if board[i-1][j:j+1] != board[i][j:j+1] {
					colS[j] = ""
				}
				// 对角线A
				if board[i-1][i-1:i] != board[i][i:i+1] {
					rowA = ""
				}
				// 对角线A
				if board[i-1][num-i:num-(i-1)] != board[i][num-(i+1):num-i] {
					rowB = ""
				}
			}

			// 判断是否还有空位
			if board[i][j:j+1] == " " {
				isFull = false
			}
		}
	}
	// 对角线A
	if rowA != "" && (rowA == "O" || rowA == "X") {
		return rowA
	}
	// 对角线B
	if rowB != "" && (rowB == "O" || rowB == "X") {
		return rowB
	}
	// 行
	for _, v := range rowS {
		if v != "" && (v == "O" || v == "X") {
			return v
		}
	}
	// 列
	for _, v := range colS {
		if v != "" && (v == "O" || v == "X") {
			return v
		}
	}
	// 如果没有选手获胜
	if isFull {
		return string("Draw")
	} else {
		return string("Pending")
	}
}

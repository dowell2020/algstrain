/*
 * @Author: dowell87
 * @Date: 2021-11-07 19:38:00
 * @Descripttion:
 * @LastEditTime: 2021-11-07 23:34:12
 */
package algs14

func tictactoe1(board []string) (res string) {
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
	}

	return string("Pending")
}

func tictactoe(board []string) (res string) {
	// 获取棋盘的大小
	isFull := false                // 0 default none space 1 exist space
	stringToNum := map[string]int{ // X:1 O:0 " ":-1000
		"X": 1,
		"O": 0,
		" ": -1000,
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 列号
		var sum = 0                       // 列和初始化
		for j := 0; j < len(board); j++ { // 行号
			sum += stringToNum[string(board[j][i])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			isFull = true
		}
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 行号
		var sum = 0                       // 行和初始化
		for j := 0; j < len(board); j++ { // 列号
			sum += stringToNum[string(board[i][j])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			isFull = true
		}
	}

	var sum1 = 0                      // 右上左下对角线之和
	var sum2 = 0                      // 左上右下对角线之和
	for i := 0; i < len(board); i++ { // 计算对角线之和
		sum1 = sum1 + stringToNum[string(board[len(board)-i-1][i])]
		sum2 = sum2 + stringToNum[string(board[i][i])]
	}
	if sum1 == len(board) || sum2 == len(board) { // winner : X
		return "X"
	} else if sum1 == 0 || sum2 == 0 { // winner : O
		return "O"
	} else if sum1 < 0 || sum2 < 0 {
		isFull = true
	}

	if isFull { // 如果不存在判断是否有空格
		return "Pending"
	}

	return "Draw"

}

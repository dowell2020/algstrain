/*
 * @Author: dowell87
 * @Date: 2021-07-08 18:57:23
 * @Descripttion:
 * @LastEditTime: 2021-07-08 22:43:32
 */
package algs1

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	newstring := ""
	n := len(s)
	cycleLen := 2*numRows - 2
	i := 0

	for i < numRows {
		for j := 0; j+i < n; j += cycleLen {
			newstring = newstring + string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				newstring = newstring + string(s[j+cycleLen-i])
			}
		}
		i++
	}
	return newstring
	// newstring := ""
	// i := 0
	// numRow := 1
	// cycleLen := 2*numRows - 2
	// for numRow <= numRows {
	// 	for i < len(s) {
	// 		l := i%cycleLen + 1
	// 		if l > numRows {
	// 			if numRow != 1 && numRow != numRows {
	// 				// fmt.Println(numRow)
	// 				if l%numRow == i%numRow-1 {

	// 					fmt.Println(numRow)
	// 					fmt.Println(i)
	// 					// newstring = newstring + string(s[i])
	// 				}
	// 				// if numRow%2 == 0 {
	// 				// 	// fmt.Println(numRow)
	// 				// 	// fmt.Println(i)
	// 				// 	// newstring = newstring + string(s[i])
	// 				// }
	// 			}
	// 		} else {
	// 			if l == numRow {
	// 				newstring = newstring + string(s[i])
	// 			}
	// 		}

	// 		i++
	// 	}
	// 	i = 0
	// 	numRow++
	// }
	// return newstring
}

// @lc code=end

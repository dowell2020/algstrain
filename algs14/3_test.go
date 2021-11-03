/*
 * @Author: dowell87
 * @Date: 2021-11-03 11:07:28
 * @Descripttion:
 * @LastEditTime: 2021-11-03 16:39:28
 */
package algs14

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	// for i := 0; i < 1; i++ {
	// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// 	Rotate(matrix)
	// 	// fmt.Println(matrix)
	// 	// fmt.Printf("a", time.Now())
	// }
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(matrix)
}

func TestIsValidSudoku(t *testing.T) {
	// for i := 0; i < 1; i++ {
	// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// 	Rotate(matrix)
	// 	// fmt.Println(matrix)
	// 	// fmt.Printf("a", time.Now())
	// }
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	res := IsValidSudoku(board)
	fmt.Println(res)
}

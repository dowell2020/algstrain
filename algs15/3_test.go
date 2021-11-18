/*
 * @Author: dowell87
 * @Date: 2021-11-18 14:27:52
 * @Descripttion:
 * @LastEditTime: 2021-11-18 14:30:34
 */
package algs15

import (
	"fmt"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	GameOfLife(board)
	fmt.Println(board)
}

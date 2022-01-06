/*
 * @Author: dowell87
 * @Date: 2021-11-07 19:38:08
 * @Descripttion:
 * @LastEditTime: 2021-11-07 22:09:49
 */
package algs14

import (
	"fmt"
	"testing"
)

func TestTictactoe(t *testing.T) {
	// // Draw
	// board := []string{"OOX", "XXO", "OXO"}

	// Pending
	// board := []string{"OOX", "XXO", "OX "}
	// O
	// board := []string{"OXX", "OOO", "XXO"}
	// X
	// board := []string{"OXX", "XXO", "OXO"}
	// X
	// board := []string{"O X", " OX", "X O"}
	// board := []string{"O X", " XO", "X O"}
	// board := []string{"   X O  O ", " X X    O ", "X  X    O ", "X    OX O ", "X   XO  O ", "X  X O  O ", "O  X O  O ", "     O  OX", "     O  O ", "   X XXXO "}
	// board := []string{"                             ", "                     O       ", "         X           X       ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "                      O      ", "                             ", "                             ", "                             ", "                             ", "                             ", "                             ", "      X                      ", "                             ", "                             ", "      O                      ", "                             ", "                             "}
	board := []string{"O X", " XO", "X O"}
	res := tictactoe(board)
	fmt.Println(res)
}

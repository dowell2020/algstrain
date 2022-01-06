/*
 * @Author: dowell87
 * @Date: 2021-10-27 17:22:44
 * @Descripttion:
 * @LastEditTime: 2021-10-27 17:50:49
 */
package algs14

import "testing"

func TestFindWords1(t *testing.T) {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	for i := 0; i < 10; i++ {
		t.Log(FindWords(board, words))
	}
}

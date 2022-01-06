/*
 * @Author: dowell87
 * @Date: 2021-10-31 12:04:43
 * @Descripttion:
 * @LastEditTime: 2021-10-31 12:16:03
 */
package algs14

import "testing"

func TestCherryPickup(t *testing.T) {
	grid := [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}

	for i := 0; i < 20; i++ {
		CherryPickup(grid)
	}
}

/*
 * @Author: dowell87
 * @Date: 2021-12-09 10:58:09
 * @Descripttion:
 * @LastEditTime: 2021-12-09 16:00:47
 */
package algs17

import (
	"fmt"
	"testing"
)

func TestClimbingLeaderboard(t *testing.T) {
	// rank := []int32{100, 100, 50, 40, 40, 20, 10}
	// score := []int32{5, 25, 50, 120}
	rank := []int32{100, 90, 90, 80, 75, 60}
	score := []int32{50, 65, 77, 90, 102}
	res := ClimbingLeaderboard(rank, score)
	fmt.Println(res)
}

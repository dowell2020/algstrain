/*
 * @Author: dowell87
 * @Date: 2021-12-09 10:56:55
 * @Descripttion:
 * @LastEditTime: 2021-12-09 16:18:50
 */
package algs17

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	rank_num := 1
	temp := 0
	score := make([]int32, len(player))
	for k, v := range player {
		rank_num = 1
		temp = 0
		for i := 0; i < len(ranked); i++ {
			if v >= ranked[i] {
				temp = 1
				break
			}
			if i < len(ranked)-1 {
				if ranked[i] != ranked[i+1] {
					rank_num++
				}
			}
		}
		if temp == 0 {
			score[k] = int32(rank_num + 1)
		} else {
			score[k] = int32(rank_num)
		}
	}
	return score
}

/*
 * @Author: dowell87
 * @Date: 2021-11-28 18:26:41
 * @Descripttion:
 * @LastEditTime: 2021-11-28 20:51:31
 */

package algs16

// 解法超时了
// func WaysToChange(n int) int {
// 	arrs := [][]int{}
// 	temparr := []int{}
// 	values := []int{25, 10, 5, 1}
// 	var dfs func(arr []int, cur int)
// 	dfs = func(arr []int, cur int) {
// 		if cur >= n {
// 			if cur == n {
// 				for j := 0; j < len(arrs); j++ {
// 					// 判断条件有问题
// 					if len(arrs[j]) == len(arr) {
// 						return
// 					}
// 				}
// 				arrs = append(arrs, arr)
// 			}
// 			return
// 		}
// 		for i := 0; i < len(values); i++ {
// 			if n-cur >= values[i] {
// 				temparr = arr
// 				dfs(append(temparr, values[i]), values[i]+cur)
// 			}
// 		}

// 	}
// 	dfs([]int{}, 0)
// 	return len(arrs) % 1000000007
// }
// 官方
// func WaysToChange(n int) int {
// 	ans := 0
// 	mod := 1000000007
// 	for i := 0; i*25 <= n; i++ {
// 		rest := n - i*25
// 		a, b := rest/10, rest%10/5
// 		ans = (ans + (a+1)*(a+b+1)%mod) % mod
// 	}
// 	return ans
// }

func WaysToChange(n int) int {
	// 假设有n+1个可能性
	dp := make([]int, n+1)
	// 零钱的集合
	coins := []int{1, 5, 10, 25}
	//刚好可以用一个硬币凑成的情况，是一种情况
	dp[0] = 1
	/**
	* dp方程：dp[i] += dp[i - coin];
	 */
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % 1000000007
		}
	}
	return dp[n]
}

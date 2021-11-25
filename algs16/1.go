/*
 * @Author: dowell87
 * @Date: 2021-11-25 16:53:51
 * @Descripttion:
 * @LastEditTime: 2021-11-25 19:20:20
 */
package algs16

/**
 * @description: 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。有效括号组合需满足：左括号必须以正确的顺序闭合。
 * @param {int} n
 * @return {*}
 */
func GenerateParenthesis(n int) []string {
	strs := []string{}
	var dfs func(i, j int, str string)
	dfs = func(i, j int, str string) {
		if 2*n == len(str) {
			strs = append(strs, str)
			return
		}
		if i > 0 {
			dfs(i-1, j, str+"(")
		}
		if i < j {
			dfs(i, j-1, str+")")
		}
	}
	dfs(n, n, "")
	return strs
}

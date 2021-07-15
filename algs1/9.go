/*
 * @Author: dowell87
 * @Date: 2021-07-15 13:00:14
 * @Descripttion:
 * @LastEditTime: 2021-07-15 13:35:12
 */
package algs1

/*
* 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
* 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*
 */

/**
 * @description: 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
 * @param {*}
 * @return {*}
 */

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 将x反向生成newx  如果newx = x
	newx := 0
	temp := x
	for x > newx {
		newx = newx*10 + temp%10
		temp = temp / 10
	}
	if newx == x {
		return true
	}
	return false
}

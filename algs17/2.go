/*
 * @Author: dowell87
 * @Date: 2021-12-09 17:19:18
 * @Descripttion:
 * @LastEditTime: 2021-12-09 18:24:42
 */
package algs17

func CanMeasureWater(x int, y int, z int) (result bool) {
	if x == z || y == z || x+y == z {
		return true
	}
	if x+y < z {
		return false
	}
	return z%GCD(x, y) == 0
}

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

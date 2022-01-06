/*
 * @Author: dowell87
 * @Date: 2021-11-11 11:45:51
 * @Descripttion:
 * @LastEditTime: 2021-11-11 17:20:44
 */
package algs15

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	num1, num2 := "123", "456"
	res := Multiply(num1, num2)
	fmt.Println(res)
}

func TestFindDiagonalOrder(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := FindDiagonalOrder(mat)
	fmt.Println(res)
}

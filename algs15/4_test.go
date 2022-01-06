/*
 * @Author: dowell87
 * @Date: 2021-11-18 15:36:16
 * @Descripttion:
 * @LastEditTime: 2021-11-18 15:52:17
 */
package algs15

import (
	"fmt"
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	num1, num2 := "1+1i", "1+1i"
	res := ComplexNumberMultiply(num1, num2)
	fmt.Println(res)
}

func TestCarPooling(t *testing.T) {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	res := CarPooling(trips, capacity)
	fmt.Println(res)
}

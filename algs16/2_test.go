/*
 * @Author: dowell87
 * @Date: 2021-11-28 18:42:14
 * @Descripttion:
 * @LastEditTime: 2021-11-28 20:29:57
 */
package algs16

import (
	"fmt"
	"testing"
)

func TestWaysToChange(t *testing.T) {
	n := 61
	res := WaysToChange(n)
	fmt.Println(res)
}

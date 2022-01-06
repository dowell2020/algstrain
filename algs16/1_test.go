/*
 * @Author: dowell87
 * @Date: 2021-11-25 16:54:06
 * @Descripttion:
 * @LastEditTime: 2021-11-25 19:20:51
 */
package algs16

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 7
	res := GenerateParenthesis(n)
	fmt.Println(res)
}

/*
 * @Author: dowell87
 * @Date: 2021-12-09 17:19:53
 * @Descripttion:
 * @LastEditTime: 2021-12-09 17:45:03
 */
package algs17

import (
	"fmt"
	"testing"
)

func TestCanMeasureWater(t *testing.T) {
	// x, y, z := 3, 5, 4
	x, y, z := 34, 5, 6
	res := CanMeasureWater(x, y, z)
	fmt.Println(res)
}

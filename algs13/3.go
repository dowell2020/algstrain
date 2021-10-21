package algs13

import (
	"fmt"
	"sort"
	"strconv"
)

/*
* 179. 最大数
*
* 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
* 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*
 */
func LargestNumber(nums []int) string {
	// s := ""
	// if len(nums) == 0 {
	// 	return s
	// }
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	fmt.Println(nums)
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

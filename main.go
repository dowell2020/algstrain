/*
 * @Author: dowell87
 * @Date: 2021-07-15 12:58:01
 * @Descripttion:
 * @LastEditTime: 2021-07-15 23:07:22
 */
package main

import (
	"fmt"

	"github.com/dowell2020/algstrain/algs1"
)

func main() {
	// res := algs1.IsPalindrome(1001)
	// fmt.Println(res)
	// res := algs1.IsMatch("aa", "a")
	// res := algs1.IsMatch("aa", "a*")
	// res := algs1.IntToRoman(332)
	// res := algs1.RomanToInt("III")
	// res := algs1.RomanToInt("IV")
	// var str = []string{"dog", "racecar", "car"}
	// var str = []string{"flower", "flow", "flight"}
	// var str = []string{"a", "b"}
	// var str = []string{"ab", "a"}
	// res := algs1.LongestCommonPrefix(str)
	var str = []int{-1, 0, 1, 2, -1, -4}
	res := algs1.ThreeSum(str)
	fmt.Println(res)
}

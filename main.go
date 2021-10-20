package main

import (
	"fmt"

	"github.com/dowell2020/algstrain/algs13"
)

func main() {

	// // FindSubstring
	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word", "good", "best", "word"}

	// res := algs12.FindSubstring(s, words)
	// nums := []int{-1, 2, 1, -4}
	// target := 1
	// res := algs13.ThreeSumClosest(nums, target)

	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	res := algs13.FourSum(nums, target)
	fmt.Println(res)
	// fmt.Println(".......")
}

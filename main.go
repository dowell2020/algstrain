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

	// nums := []int{1, 0, -1, 0, -2, 2}
	// target := 0
	// res := algs13.FourSum(nums, target)
	// nums := []int{3, 30, 34, 5, 9}
	// res := algs13.LargestNumber(nums)
	// fmt.Println(res)

	// nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	// res := algs13.FindKthLargest(nums, 4)
	// nums := []int{3, 6, 9, 1}
	// res := algs13.MaximumGap(nums)
	// fmt.Println(res)

	// n := 6
	// speed := []int{2, 10, 3, 1, 5, 8}
	// efficiency := []int{5, 4, 3, 9, 7, 2}
	// k := 2
	// res := algs13.MaxPerformance(n, speed, efficiency, k)
	// fmt.Println(res)
	box := [][]int{{1, 1, 1}, {2, 3, 4}, {2, 3, 4}, {2, 6, 7}, {3, 4, 5}}
	res := algs13.PileBox(box)
	fmt.Println(res)

	// fmt.Println(".......")
}

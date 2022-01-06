package algs13

import "fmt"

func MaxPerformance(n int, speed []int, efficiency []int, k int) int {

	res := []int{}

	for i := 0; i < n; i++ {
		res = append(res, speed[i]*efficiency[i])
	}
	fmt.Println(res)
	return 0
}

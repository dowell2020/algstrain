package algs13

import (
	"sort"
)

// func PileBox(box [][]int) int {

// 	// 计算的高度
// 	// 最后一层的深度和高度
// 	k := 0
// 	for i := 1; i < len(box); i++ {
// 		k = 0
// 		// 当前盒子能放在那个盒子上面
// 		for j := 0; j < len(sum); j++ {
// 			// 能堆叠
// 			if sum[j][1] < box[i][0] && sum[j][2] < box[i][1] && sum[j][3] < box[i][2] {
// 				sum[j][0] = sum[j][0] + box[i][2]
// 				sum[j][1] = box[i][0]
// 				sum[j][2] = box[i][1]
// 				sum[j][3] = box[i][2]
// 				k++
// 			}
// 		}
// 		// 不能堆叠,添加新的一堆
// 		if k == 0 {
// 			tempsum := []int{box[i][2], box[i][0], box[i][1], box[i][2]}
// 			for z := i - 1; z >= 0; z-- {
// 				// 能添加多少层的高度计算
// 				if box[z][0] < tempsum[1] && box[z][1] < tempsum[2] && box[z][2] < tempsum[3] {
// 					tempsum[0] = tempsum[0] + box[z][2]
// 					tempsum[1] = box[z][0]
// 					tempsum[2] = box[z][1]
// 					tempsum[3] = box[z][2]
// 				}
// 			}
// 			tempsum[1] = box[i][0]
// 			tempsum[2] = box[i][1]
// 			tempsum[3] = box[i][2]
// 			fmt.Print(i)
// 			fmt.Print(tempsum)

// 			sum = append(sum, tempsum)
// 		}
// 	}
// 	sort.Slice(sum, func(i, j int) bool {
// 		return sum[i][0] > sum[j][0]
// 	})
// 	return sum[0][0]
// }

func PileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	// 初始化多少堆
	sum := make([]int, len(box))
	for i := range sum {
		sum[i] = box[i][2]
	}
	for i := 0; i < len(box); i++ {
		temp := 0
		for j := 0; j < i; j++ {
			// 将后面的箱子往上面堆
			if sum[j] > temp && box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				temp = sum[j]
			}
		}
		// 如果能对上去，高度++
		sum[i] = sum[i] + temp
	}
	// 对高度内容排序,获取最大的内容
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	return sum[0]
}

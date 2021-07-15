/*
 * @Author: dowell87
 * @Date: 2021-07-15 16:12:30
 * @Descripttion:
 * @LastEditTime: 2021-07-15 18:19:13
 */
package algs1

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func IntToRoman(num int) string {
	var romanS = []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var romanI = []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var str = ""
	// var TempInt int
	// 一直减
	i := len(romanI) - 1
	for num != 0 {
		if num-romanI[i] >= 0 {
			num -= romanI[i]
			str = str + romanS[i]
		} else {
			i--
		}
	}

	return str
}

// @lc code=end

// 遇到的问题
//  赋值完成后不是顺序循环输出的
// roman := map[int]string{
// 		1000: "M",
// 		900:  "CM",
// 		500:  "D",
// 		400:  "CD",
// 		100:  "C",
// 		90:   "XC",
// 		50:   "L",
// 		40:   "XL",
// 		10:   "X",
// 		9:    "IX",
// 		5:    "V",
// 		4:    "IV",
// 		1:    "I",
// 	}

// func IntToRoman(num int) string {
// 	romanS := map[int]string{1: "M", 2: "CM", 3: "D", 4: "CD", 5: "C", 6: "XC", 7: "L", 8: "XL", 9: "X", 10: "IX", 11: "V", 12: "IV", 13: "I"}
// 	romanI := map[int]int{1: 1000, 2: 900, 3: 500, 4: 400, 5: 100, 6: 90, 7: 50, 8: 40, 9: 10, 10: 9, 11: 5, 12: 4, 13: 1}
// 	str := ""
// 	for k, v := range romanI {
// 		if num/v >= 1 {
// 			// fmt.Println(num / v)
// 			// fmt.Println(v)
//         	// 当有打印输出的时候for循环的内容可能有变化
// 			num = num % v
// 			// fmt.Println(num)
// 			str = str + romanS[k]
// 		}
// 	}
// 	return str
// }

// 贪心算法

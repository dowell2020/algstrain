/*
 * @Author: dowell87
 * @Date: 2021-11-21 17:41:52
 * @Descripttion:
 * @LastEditTime: 2021-11-21 21:31:34
 */
package algs15

// func ConcatenatedBinary(n int) int {
// 	strs := ""
// 	for i := 1; i <= n; i++ {
// 		strs = strs + strconv.FormatInt(int64(i), 2)
// 	}
// 	fmt.Println(strs)
// 	// 直接将字符传转为10进入会超出限制
// 	num, _ := strconv.Atoi(strs)
// 	fmt.Println(num)
// 	return int(num % (1e9 + 7))
// }

/*
* 网上寻找的答案
*
 */

func ConcatenatedBinary(n int) int {
	length, res := 0, 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			length++
		}
		res = res<<length | i
	}
	return res % (1e9 + 7)
}

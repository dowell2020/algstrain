package algs12

import (
	"strconv"
	"strings"
)

/*
*  解题思路:
*
*
*
*
*
 */

func RestoreIpAddresses(s string) []string {
	ips := []string{}
	if len(s) < 4 || len(s) > 12 {
		return ips
	}
	// ipAddr := []string{}
	// 每一位不能大于255,不能以0开头 长度为3位
	var dfs func(ipAddr []string, str string)

	dfs = func(ipAddr []string, str string) {
		// 结束的状态
		lenAddr := len(ipAddr)
		if lenAddr == 4 && len(str) == 0 {
			ips = append(ips, strings.Join(ipAddr, "."))
			return
		}
		// 1个字符
		for i := 0; i < 3; i++ {
			// ipAddr = append(ipAddr, str[0:i])
			// fmt.Println("aaa")
			// fmt.Println(4 - lenAddr)
			// fmt.Println("bbb")
			if len(str)-(i+1) >= 3-lenAddr && len(str)-(i+1) <= (3-lenAddr)*3 {
				// 对数字的判断,大于255的情况
				if i == 2 {
					tempnum, _ := strconv.Atoi(str[0 : i+1])
					if tempnum > 255 {
						break
					}
				}
				// 多位数字前面有0的情况
				if i >= 1 {
					tempnum, _ := strconv.Atoi(str[:1])
					if tempnum == 0 {
						break
					}
				}

				tempAddr := append(ipAddr, str[0:i+1])
				tempstr := str[i+1 : len(str)]
				dfs(tempAddr, tempstr)
			}
			// 剩余字符串大于
			// dfs(ipAddr, str[i:len(str)])
			// fmt.Println(ipAddr)
		}
	}
	dfs(ips, s)
	return ips

}

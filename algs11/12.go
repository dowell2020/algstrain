package algs11

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	pre := strs[0]
	for _, k := range strs {
		for strings.Index(k, pre) != 0 {
			if len(pre) == 0 {
				return ""
			}
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}

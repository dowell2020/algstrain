package algs12

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}
	var resMap = make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		var curStrArr = strings.Split(strs[i], "")
		sort.Strings(curStrArr)
		var curStrIndex = strings.Join(curStrArr, "")
		resMap[curStrIndex] = append(resMap[curStrIndex], strs[i])
	}
	for _, value := range resMap {
		res = append(res, value)
	}
	return res
}

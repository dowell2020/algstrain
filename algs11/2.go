package algs11

func LengthOfLongestSubstring(s string) int {
	cot := len(s)
	if cot == 0 {
		return 0
	}
	tempMap := make(map[uint8]int)
	l, r, maxLen := 0, 1, 1
	tempMap[s[l]] = l
	for l < r && r < cot {
		val, ok := tempMap[s[r]]
		if ok && val >= l {
			l = val + 1
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}

		tempMap[s[r]] = r
		r++
	}

	return maxLen
}

package algs11

func LongestPalindrome(s string) string {
	var stringlen int = len(s)
	if stringlen <= 1 {
		return s
	}
	var i int = 0
	var maxLen = 0
	var res = ""
	for i < stringlen {
		var pointerL = i
		var pointerR = i
		for pointerL >= 0 && s[pointerL] == s[i] {
			pointerL--
		}
		for pointerR < stringlen && s[pointerR] == s[i] {
			pointerR++
		}
		var nextPoint = pointerR
		for pointerL >= 0 && pointerR < stringlen && s[pointerR] == s[pointerL] {
			pointerL--
			pointerR++
		}
		var tmpMaxLen = pointerR - pointerL - 1
		if tmpMaxLen > maxLen {
			res = s[pointerL+1 : pointerR]
			maxLen = tmpMaxLen
		}
		i = nextPoint
	}
	return res
}

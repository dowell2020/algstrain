package algs12

func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 || n == 0 {
		return false
	}
	// 另一半
	akuo := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	kuo := []byte{}
	for i := 0; i < n; i++ {
		if akuo[s[i]] > 0 {
			if len(kuo) == 0 || kuo[len(kuo)-1] != akuo[s[i]] {
				return false
			}
			kuo = kuo[:len(kuo)-1]
		} else {
			kuo = append(kuo, s[i])
		}
	}
	return len(kuo) == 0
}

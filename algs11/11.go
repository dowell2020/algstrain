package algs11

func RomanToInt(s string) int {
	var mp = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var num int

	tmp := 1000
	for i := 0; i < len(s); i++ {
		now := mp[s[i]]
		if tmp < now {
			num -= 2 * tmp
		}
		num += now
		tmp = now
	}
	return num
}

package algs11

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

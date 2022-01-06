package algs11

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 将x反向生成newx  如果newx = x
	newx := 0
	temp := x
	for x > newx {
		newx = newx*10 + temp%10
		temp = temp / 10
	}
	if newx == x {
		return true
	}
	return false
}

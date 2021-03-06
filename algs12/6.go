package algs12

func CanJump(nums []int) bool {
	max := 0
	for k, n := range nums {
		if n == 0 && k != len(nums)-1 {
			if max <= 1 {
				return false
			}
		}
		max = max - 1
		if n > max {
			max = n
		}

	}
	return true
}

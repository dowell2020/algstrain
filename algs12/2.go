package algs12

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n > 2 {
		m := len(nums)
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if nums[i] == nums[j] {
					nums = append(nums[:i], nums[i+1:]...)
					m--
					j--
				}
			}
		}
	} else if n == 2 {
		if nums[0] == nums[1] {
			nums = nums[:1]
		}
	}

	return len(nums)
}

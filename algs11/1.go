package algs11

func TwoSum(nums []int, target int) []int {
	newnums := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				newnums[0] = i
				newnums[1] = j
			}
		}
	}
	return newnums
}

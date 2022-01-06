package algs12

func permute(nums []int) [][]int {
	res := [][]int{}
	backPermute(0, nums, &res)
	return res
}

func backPermute(first int, nums []int, res *[][]int) {
	if first == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}
	m := map[int]bool{}
	for i := first; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[first], nums[i] = nums[i], nums[first]
		backPermute(first+1, nums, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

package algs12

func Subsets(nums []int) [][]int {
	sums := [][]int{}
	// 临时子集
	tempnum := []int{}
	// 列举出所有可能性
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			sums = append(sums, append([]int(nil), tempnum...))
			return
		}
		tempnum = append(tempnum, nums[cur])
		dfs(cur + 1)
		tempnum = tempnum[:len(tempnum)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return sums

}

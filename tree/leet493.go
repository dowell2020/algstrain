package tree

type PairsTree struct {
	Val   int
	Sum   int
	Left  *PairsTree
	Right *PairsTree
}

// 创建树
func (tree *PairsTree) CreatPairsTree(val int) *PairsTree {
	if tree == nil {
		return &PairsTree{val, 1, nil, nil}
	}
	if tree.Val == val {
		tree.Sum++
	} else if tree.Val < val {
		tree.Right = tree.Right.CreatPairsTree(val)
	} else if tree.Val > val {
		tree.Sum++
		tree.Left = tree.Left.CreatPairsTree(val)
	}
	return tree
}

// 统计
func (tree *PairsTree) SumPairsTree(cur int) int {
	if tree == nil {
		return 0
	}
	if tree.Val >= cur {
		return tree.Left.SumPairsTree(cur)
	} else {
		return tree.Sum + tree.Right.SumPairsTree(cur)
	}

}

func reversePairs(nums []int) int {
	ans, n := 0, len(nums)
	if n > 0 {
		tree := &PairsTree{nums[n-1], 1, nil, nil}
		for i := n - 2; i >= 0; i-- {
			ans += tree.SumPairsTree((nums[i] + 1) >> 1)
			tree.CreatPairsTree(nums[i])
		}
	}
	return ans
}

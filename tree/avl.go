package tree

import (
	"log"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type AvlNode struct {
	Val   int
	High  int
	Left  *AvlNode
	Right *AvlNode
}

// 获取高度
func GetHigh(avl *AvlNode) int {
	if avl == nil {
		return 0
	}
	return avl.High
}

// 获取较大值

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FindMax 查找最大节点

func (avl *AvlNode) FindMax() *AvlNode {
	if avl.Right != nil {
		return avl.Right.FindMax()
	}
	return avl
}

// FindMin 查找最小节点
func (avl *AvlNode) FindMin() *AvlNode {
	if avl.Left != nil {
		return avl.Left.FindMax()
	}
	return avl
}

//GetBF 获取节点左右子树高度差绝对值
//将二叉树上节点的左子树高度减去右子树高度取绝对值(Balance Factor)
func (avl *AvlNode) GetBF() int {
	var lh, rh int
	if avl.Left != nil {
		lh = GetHigh(avl.Left)
	}
	if avl.Right != nil {
		rh = GetHigh(avl.Right)
	}
	bf := lh - rh
	if bf < 0 {
		bf = bf * -1
	}
	return bf
}

//LeftRotation 左旋
//a为最小失衡子数的根节点
//问题：在右子树插入右孩子导致AVL失衡
//return 新的平衡树的根节点

func (avl *AvlNode) LeftRotation() *AvlNode {
	tmpnode := avl.Right
	avl.Right = tmpnode.Left
	tmpnode.Left = avl
	avl.High = Max(GetHigh(avl.Left), GetHigh(avl.Right)) + 1
	tmpnode.High = Max(GetHigh(avl.Left), GetHigh(avl.Right)) + 1
	return tmpnode
}

//RightRotation 右旋
//a为最小失衡树的根节点
//问题：在左子树上插入左孩子导致AVL树失衡
//return 新的平衡树的根节点

func (avl *AvlNode) RightRotation() *AvlNode {
	tmpnode := avl.Left
	avl.Left = tmpnode.Right
	tmpnode.Right = avl
	avl.High = Max(GetHigh(avl.Left), GetHigh(avl.Right)) + 1
	tmpnode.High = Max(GetHigh(tmpnode.Left), GetHigh(tmpnode.Right)) + 1
	return tmpnode
}

//RightLeftRotation  右左双旋转
//问题：通常因为在右子树上插入左孩子导致AVL失衡
//解发：先右旋后左旋调整
//return 新的平衡树根节点
func (avl *AvlNode) RightLeftRotation() *AvlNode {
	avl.Right = avl.Right.RightRotation()
	return avl.LeftRotation()
}

//LeftRightRotation  左右双选择
//问题：通常因为在左子树上插入右孩子导致AVL失衡
//解发：先左旋后右旋调整
//return 新的平衡树根节点
func (avl *AvlNode) LeftRightRotation() *AvlNode {
	avl.Left = avl.Left.LeftRotation()
	return avl.RightRotation()
}

//InsertAVL avl树插入操作
//递归
// 1个节点的高度初始值为1
//返回插入后的根节点
func (avl *AvlNode) InsertAvlTree(v int) *AvlNode {
	if avl == nil { //空树，插入第一个节点
		avl = &AvlNode{v, 0, nil, nil}
	} else if v < avl.Val { //待插入的值比单前节点值小，往左子树插入
		avl.Left = avl.Left.InsertAvlTree(v)
		if GetHigh(avl.Left)-GetHigh(avl.Right) == 2 { //插入新节点后avl树失衡
			if v < avl.Left.Val { //情况1：v插入到左子树的左孩子节点，只需要进行一次右旋转
				avl = avl.RightRotation()
			} else if v > avl.Left.Val { //情况2：v插入到左子树的右孩子节点，需要先左旋转再右旋转
				avl = avl.LeftRightRotation()
			}
		}
	} else if v > avl.Val {
		avl.Right = avl.Right.InsertAvlTree(v)
		if GetHigh(avl.Right)-GetHigh(avl.Left) == 2 { //在右子树插入新节点后avl树失衡
			if v > avl.Right.Val { //情况4：v插入到右子树的右孩子节点，只需要进行一次左旋转
				avl = avl.LeftRotation()
			} else if v < avl.Right.Val { //情况3：v插入到右子树的左孩子节点，需要先右旋转再左旋转
				avl = avl.RightLeftRotation()
			}
		}
	}
	avl.High = Max(GetHigh(avl.Left), GetHigh(avl.Right)) + 1
	return avl
}

//DeleteAVL 删除avl树中值为v的节点
//维护二叉排序树规则
//平衡avl二叉树

func (tree *AvlNode) DeleteAVL(v int) *AvlNode {
	if tree == nil {
		return nil
	}
	if tree.Val == v { //tree为待删除的节点
		//删除后维护成新的二叉排序树
		if tree.Left != nil && tree.Right != nil { //若待删除节点同时存在左右子树
			if GetHigh(tree.Left) > GetHigh(tree.Right) { //左子树高于右子树，则取左子树最大值取代被删除节点
				tmp := tree.Left.FindMax()
				tree.Val = tmp.Val
				tree.Left = tree.Left.DeleteAVL(tmp.Val)
			} else { //右子树高于左子树，则取右子树最小值取代被删除节点
				tmp := tree.Right.FindMin()
				tree.Val = tmp.Val
				tree.Right = tree.Right.DeleteAVL(tmp.Val)
			}
		} else {
			if tree.Left != nil {
				tree = tree.Left
			}
			if tree.Right != nil {
				tree = tree.Right
			}
		}
		return tree
	} else if tree.Val < v { //待删除节点比当前节点大
		tree.Right = tree.Right.DeleteAVL(v)
		if GetHigh(tree.Left)-GetHigh(tree.Right) == 2 { //删除节点后avl树失衡
			if v < tree.Left.Val { //情况1：左子树的左孩子节点过高，只需要进行一次右旋转
				tree = tree.RightRotation()
			} else if v > tree.Left.Val { //情况2：左子树的右孩子节点过高，需要先左旋转再右旋转
				tree = tree.LeftRightRotation()
			}
		}
	} else if tree.Val > v {
		tree.Left = tree.Left.DeleteAVL(v)
		if GetHigh(tree.Right)-GetHigh(tree.Left) == 2 { //在右子树插入新节点后avl树失衡
			if v > tree.Right.Val { //情况4：右子树的右孩子节点过高，只需要进行一次左旋转
				tree = tree.LeftRotation()
			} else if v < tree.Right.Val { //情况3：右子树的左孩子节点过高，需要先右旋转再左旋转
				tree = tree.RightLeftRotation()
			}
		}
	}
	return tree
}

func (avl *AvlNode) PrintSortTree() {
	if avl == nil {
		return
	}
	avl.Left.PrintSortTree()
	log.Println(avl.Val)
	avl.Right.PrintSortTree()
}

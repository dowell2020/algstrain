package tree

import (
	"errors"
	"log"
)

type BstNode struct {
	Val   string
	Data  string
	Left  *BstNode
	Right *BstNode
}

// 插入
func (bst *BstNode) BstInsert(val, data string) error {
	if bst == nil {
		return errors.New("不能在空对象对插入数组")
	}
	switch {
	case bst.Val == val:
		return nil
	case bst.Val > val:
		if bst.Left == nil {
			bst.Left = &BstNode{Val: val, Data: data}
			return nil
		}
		return bst.Left.BstInsert(val, data)
	case bst.Val < val:
		if bst.Right == nil {
			bst.Right = &BstNode{Val: val, Data: data}
			return nil
		}
		return bst.Right.BstInsert(val, data)
	}
	return nil
}

// 查找

func (bst *BstNode) BstFind(val string) (string, bool) {
	if bst == nil {
		log.Fatalln("不能再空对象查找数组")
		return "", false
	}
	switch {
	case bst.Val == val:
		return bst.Data, true
	case bst.Val > val:
		return bst.Left.BstFind(val)
	default:
		return bst.Right.BstFind(val)
	}
}

// 查找最大值

func (bst *BstNode) BstFindMax(parent *BstNode) (*BstNode, *BstNode) {
	if bst.Right == nil {
		return bst, parent
	}
	return bst.Right.BstFindMax(bst)
}

// 替换
func (bst *BstNode) BstReplaceNode(parent, node *BstNode) error {
	if bst == nil {
		return errors.New("未找到替换节点")
	}
	if bst == parent.Left {
		parent.Left = node
		return nil
	}
	parent.Right = node
	return nil
}

// 删除

func (bst *BstNode) BstDeleteNode(val string, parent *BstNode) error {
	if bst == nil {
		return errors.New("未找到删除节点")
	}
	switch {
	case bst.Val < val:
		return bst.Left.BstDeleteNode(val, bst)
	case bst.Val > val:
		return bst.Right.BstDeleteNode(val, bst)
	default:
		// 如果要删除的节点是叶子节点
		if bst.Left == nil && bst.Right == nil {
			bst.BstReplaceNode(parent, nil)
			return nil
		}
		// 如果要删除的节点左孩子节点为nil
		if bst.Left == nil {
			bst.BstReplaceNode(parent, bst.Right)
			return nil
		}
		// 如果要删除的节点右孩子节点为nil
		if bst.Right == nil {
			bst.BstReplaceNode(parent, bst.Left)
			return nil
		}
		// 如果有两个孩子节点，在左子树中找出最大的值
		replacement, replParent := bst.Left.BstFindMax(bst)
		// 用replacement的数值替代n节点的值域
		bst.Val = replacement.Val
		bst.Data = replacement.Data
		// 删除这个替代节点
		return replacement.BstDeleteNode(replacement.Val, replParent)
	}
}

// 遍历

func (bst *BstNode) BstLeftTraverse(f func(*BstNode)) {
	if bst == nil {
		return
	}
	f(bst)
	bst.Left.BstLeftTraverse(f)
	bst.Right.BstLeftTraverse(f)
}

func (bst *BstNode) BstMidTraverse(f func(*BstNode)) {
	if bst == nil {
		return
	}
	bst.Left.BstMidTraverse(f)
	f(bst)
	bst.Right.BstMidTraverse(f)
}

func (bst *BstNode) BstRightTraverse(f func(*BstNode)) {
	if bst == nil {
		return
	}
	bst.Left.BstRightTraverse(f)
	bst.Right.BstRightTraverse(f)
	f(bst)
}

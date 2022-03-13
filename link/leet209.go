package link

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// 提交还是没有通过

// func insert1(aNode *Node, x int) *Node {
// 	if aNode == nil {
// 		aNode = &Node{x, nil}
// 		return aNode
// 	}
// 	aNode.Next = insert1(aNode.Next, x)
// 	return aNode
// }

// 考虑链表为循环链表
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode := &Node{Val: x}
		aNode.Next = aNode
		return aNode
	}
	// 找到真正的起点
	if aNode.Val <= aNode.Next.Val {
		aNode.Next = &Node{Val: x, Next: aNode.Next}
	}
	// 找到插入位置
	if aNode.Next.Val < x {
		aNode.Next = &Node{Val: x, Next: aNode.Next}
	}
	return aNode
}

func (nd *Node) ViewNode() {
	if nd == nil {
		fmt.Println("-> 空链表!")
		return
	}
	fmt.Printf("%d -> ", nd.Val)
	if nd.Next != nil {
		nd.Next.ViewNode()
	}
}

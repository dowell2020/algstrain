package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
	}
}
func (n *ListNode) addNode(v int) {
	if n.Next == nil {
		n.Next = NewListNode(v)
	} else {
		n.Next.addNode(v)
	}
}

func (n *ListNode) ViewNode() {
	if n == nil {
		fmt.Println("-> 空链表!")
		return
	}
	fmt.Printf("%d -> ", n.Val)
	if n.Next != nil {
		n.Next.ViewNode()
	}
}

func (n *ListNode) deleteNode(node *ListNode) {
	if n.Val == node.Val {
		*n = *n.Next
	} else {
		n.Next.deleteNode(node)
	}
}

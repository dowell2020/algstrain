package link

import "fmt"

// 定义节点
type LinkNode struct {
	Value int
	Next  *LinkNode
}

func NewLinkNode(v int) *LinkNode {
	return &LinkNode{v, nil}
}

// 添加节点
func (ln *LinkNode) addNode(v int) {
	if ln.Next == nil {
		ln.Next = &LinkNode{v, nil}
	} else {
		ln.Next.addNode(v)
	}
}

// 遍历链表
func (ln *LinkNode) viewlink() {
	if ln == nil {
		fmt.Println("-> 空链表!")
		return
	}
	fmt.Printf("%d -> ", ln.Value)
	if ln.Next != nil {
		ln.Next.viewlink()
	}
}

// 查找节点
func (ln *LinkNode) SearchNode(v int) bool {
	if ln.Value == v {
		return true
	}
	ln.Next.SearchNode(v)
	return false
}

// 获取链表长度
func (ln *LinkNode) Count(sum int) int {

	if ln.Next != nil {
		sum = sum + 1
		ln.Next.Count(sum)
	}
	return sum
}

func (ln *LinkNode) deleteNode(node *ListNode) {
	*node = *node.Next
}

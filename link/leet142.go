package link

func detectCycle(head *ListNode) *ListNode {
	// 判断只有第一个值的时候
	if head == nil || head.Next == nil {
		return nil
	}
	// 循环1
	node := head
	for node.Next != nil {
		if node.Next == head {
			// 证明是个环
			return head
		}
	}
	return detectCycle(head.Next)
}

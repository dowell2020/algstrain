package link

func nextLargerNodes(head *ListNode) (data []int) {
	if head == nil {
		return []int{0}
	}
	next := head.Next
	for next != nil {
		if next.Val > head.Val {
			data = append(data, next.Val)
			break
		}
		if next.Next == nil {
			data = append(data, 0)
			break
		} else {
			next = next.Next
		}
	}
	data = append(data, nextLargerNodes(head.Next)...)
	return data
}

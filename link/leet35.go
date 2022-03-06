package link

import "fmt"

type RandLink struct {
	Val    int
	Next   *RandLink
	Random *RandLink
}

// 插入数据
func (rl *RandLink) AppendRandomList(k, v int) {
	if rl.Next == nil {
		if v >= 0 {
			rl.Next = &RandLink{k, nil, &RandLink{v, nil, nil}}
		} else {
			rl.Next = &RandLink{k, nil, nil}
		}
	} else {
		rl.Next.AppendRandomList(k, v)
	}
}

// 遍历数据
func (rl *RandLink) TravelRandomList() {
	if rl == nil {
		fmt.Println("-> 空链表!")
		return
	}
	fmt.Println(rl.Val)
	if rl.Random != nil {
		fmt.Println(rl.Random.Val)
	}
	fmt.Println("..........................")
	if rl.Next != nil {
		rl.Next.TravelRandomList()
	}
}

func copyRandomList2(head *RandLink) (node *RandLink) {
	if head == nil {
		return nil
	}
	node.Val = head.Val
	node.Random = head.Random
	if head.Next != nil {
		node.Next = copyRandomList(head.Next)
	}
	return node
}

// 正确的
func copyRandomList1(head *RandLink) (node *RandLink) {
	if head == nil {
		return nil
	}
	node = &RandLink{head.Val, nil, nil}
	if head.Next != nil {
		node.Next = copyRandomList(head.Next)
	}
	if head.Random != nil {
		node.Random = copyRandomList(head.Random)
	}
	return node
}

func copyRandomList(head *RandLink) (node *RandLink) {
	maps := make(map[*RandLink]*RandLink)
	cur := head
	for cur != nil {
		temp := &RandLink{
			Val: cur.Val,
		}
		maps[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		node.Next = maps[cur.Next]
		node.Random = maps[cur.Random]
		cur = cur.Next
	}
	return node
}

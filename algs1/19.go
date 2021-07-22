/*
 * @Author: dowell87
 * @Date: 2021-07-22 22:56:16
 * @Descripttion:
 * @LastEditTime: 2021-07-22 23:17:25
 */
package algs1

/**
 * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
*/

// 链表这一块确实不太会
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next

}

package link

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// 例子3
	node := NewListNode(1)
	res := detectCycle(node)
	fmt.Println(res)
	// 例子2
	node.addNode(2)
	node.Next = node
	res = detectCycle(node)
	fmt.Println(res)
}

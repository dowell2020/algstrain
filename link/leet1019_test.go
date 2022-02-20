package link

import (
	"fmt"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	// node := NewListNode(2)
	// node.addNode(1)
	// node.addNode(5)
	// node.ViewNode()
	node := NewListNode(2)
	node.addNode(7)
	node.addNode(4)
	node.addNode(3)
	node.addNode(5)
	res := nextLargerNodes(node)
	fmt.Println(res)
}

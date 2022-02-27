package link

import "testing"

func TestDeleteNode(t *testing.T) {
	node := NewListNode(4)
	node.addNode(5)
	node.addNode(9)
	node.addNode(1)
	// node.ViewNode()s t
	delnode := NewListNode(5)
	node.deleteNode(delnode)
	node.ViewNode()
}

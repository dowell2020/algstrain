package link

import (
	"testing"
)

func TestInsertLink(t *testing.T) {
	// 文件助手
	var LK *Node
	// LK := &Node{1, nil}
	// fmt.Println(LK)
	LK = insert(LK, 3)
	LK = insert(LK, 4)
	LK = insert(LK, 1)
	LK = insert(LK, 2)
	// nd = insert(nd, 3)
	// nd = insert(nd, 4)
	LK.ViewNode()
}

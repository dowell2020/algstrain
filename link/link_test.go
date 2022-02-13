package link

import (
	"fmt"
	"testing"
)

func TestLinkNode(t *testing.T) {
	tlink := NewLinkNode(3)
	// 添加节点
	tlink.addNode(1)
	tlink.addNode(-1)
	// 再次遍历
	tlink.viewlink()
	// 添加更多节点
	tlink.addNode(10)
	tlink.addNode(45)
	tlink.addNode(30)
	tlink.addNode(5)
	tlink.addNode(20)
	// 添加已存在节点
	tlink.addNode(5)
	// 再次遍历
	tlink.viewlink()

	if tlink.SearchNode(5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}
}

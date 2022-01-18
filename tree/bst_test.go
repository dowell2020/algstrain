package tree

import (
	"fmt"
	"testing"
)

func TestBstTree(t *testing.T) {
	// var bst *BstNode
	bst := &BstNode{"asd", "asd", nil, nil}
	fmt.Print("Sorted values: | ")
	arr := []string{
		"asd1",
		"asd22",
		"asd333",
		"asd4444",
		"asd55555",
		"asd666666",
		"asd7777777",
		"asd88888888",
		"asd999999999",
		"asd1010101010",
		"asd11111111111",
		"asd121212121212",
		"asd1313131313131",
		"asd14141414141414",
	}
	for _, v := range arr {
		err := bst.BstInsert(v, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	// 匿名函数简化操作
	bst.BstLeftTraverse(func(node *BstNode) { fmt.Print("Val: ", node.Val, "data: ", node.Data, " | ") })
	fmt.Println("------")
	bst.BstMidTraverse(func(node *BstNode) { fmt.Print("Val: ", node.Val, "data: ", node.Data, " | ") })
	fmt.Println("------")
	bst.BstRightTraverse(func(node *BstNode) { fmt.Print("Val: ", node.Val, "data: ", node.Data, " | ") })
	// fmt.Println(bst)
}

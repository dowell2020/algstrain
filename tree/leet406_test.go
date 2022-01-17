package tree

import (
	"fmt"
	"log"
	"testing"
)

func TestInorderSuccessor(t *testing.T) {
	// p := 1
	// root := []int{2, 1, 3}
	p := 6
	root := []int{5, 3, 6, 2, 4, 1}
	newp := NewTreeNode(p)
	var node *TreeNode
	for _, v := range root {
		node = node.CreateTreeNode(v)
		log.Println("root=> key:", node.Val)
	}
	// fmt.Println(node.Val)
	// fmt.Println(node.Left.Val)
	// fmt.Println(node.Right.Val)
	// node.FrontTreeNode()
	// node.MidTreeNode()
	// node.BackTreeNode()
	res := inorderSuccessor(node, newp)
	fmt.Println(res)

}

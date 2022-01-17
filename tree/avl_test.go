package tree

import (
	"log"
	"testing"
)

func TestAvlTree(t *testing.T) {

	// array := []int{2, 1, 3}
	array := []int{5, 3, 1, 8, 9, 10, 11, 2, 4, 7, 6, 12}
	var avlnode *AvlNode
	for _, v := range array {
		avlnode = avlnode.InsertAvlTree(v)
		log.Println("root=> key:", avlnode.Val, "high:", avlnode.High)
	}
	log.Println("-----------------")
	avlnode.PrintSortTree()
}

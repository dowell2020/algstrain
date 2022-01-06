package tree

import (
	"fmt"
	"testing"
)

func TestRemoveSubfolders(t *testing.T) {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	// folder := []string{"/a", "/a/b/c", "/a/b/d"}
	// folder := []string{"/a/b/c", "/a/b/d", "/a/b/ca"}

	res := removeSubfolders(folder)
	fmt.Println(res)
}

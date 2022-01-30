package strings

import (
	"fmt"
	"testing"
)

func TestCountValidWords(t *testing.T) {
	str := "cat and  dog"
	// str := "!this  1-s b8d!"
	res := countValidWords(str)
	fmt.Println(res)
}

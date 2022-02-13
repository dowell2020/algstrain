package link

import (
	"fmt"
	"testing"
)

func TestCountValidWords(t *testing.T) {
	str := "cat and  dog"
	res := countValidWords(str)
	fmt.Println(res)
}

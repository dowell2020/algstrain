package link

import (
	"fmt"
	"testing"
)

func TestRandomList(t *testing.T) {

	var rk RandLink
	// randlink := NewRandLink(3)
	// 插入数据
	rk.AppendRandomList(7, -1)
	rk.AppendRandomList(13, 0)
	rk.AppendRandomList(11, 4)
	rk.AppendRandomList(10, 2)
	rk.AppendRandomList(1, 0)

	// 遍历数据
	rk.TravelRandomList()
	fmt.Println("++++++++++++++++")
	newrk := copyRandomList(&rk)
	newrk.TravelRandomList()
}

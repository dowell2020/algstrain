package link

import (
	"testing"
)

func TestCirCular(t *testing.T) {
	clink := NewCirLink()
	// 5,7,10,13,100,3000
	clink.AppendCirLink(5)
	clink.AppendCirLink(7)
	clink.AppendCirLink(10)
	clink.AppendCirLink(13)
	clink.AppendCirLink(100)
	clink.AppendCirLink(3000)
	clink.TraversalCirLink()
	clink.DeleteCirLink(13)
	clink.TraversalCirLink()
}

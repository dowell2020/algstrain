package link

import (
	"errors"
	"fmt"
)

type CirLink struct {
	Val  interface{}
	Next *CirLink
}

func NewCirLink() *CirLink {
	return &CirLink{}
}

// 插入值
func (c *CirLink) AppendCirLink(data interface{}) {
	if c.Next == nil {
		c.Next = &CirLink{data, c}
		return
	}
	node := c.Next
	if node.Next.Val == nil {
		node.Next = &CirLink{data, node.Next}
		return
	}
	c.Next.AppendCirLink(data)
}

// 删除值

func (c *CirLink) DeleteCirLink(data interface{}) error {
	node := c.Next
	if node.Val == data {
		c.Next = node.Next
		return nil
	}
	if node.Next.Val == nil {
		err := errors.New("未找到对应值")
		return err
	}
	c.Next.DeleteCirLink(data)
	return nil
}

// 遍历数值

func (c *CirLink) TraversalCirLink() {
	node := c.Next
	fmt.Printf("遍历信息为:val=%d\n", node.Val)
	if node.Next.Val != nil {
		node.TraversalCirLink()
	}
}

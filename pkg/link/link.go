package link

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func NewLink(val int) *LinkNode {
	return &LinkNode{
		Value: val,
	}
}

func (link *LinkNode) Add(val int) {
	NewLink(val)
	point := link

	for point.Next != nil {
		point = point.Next
	}
	newNode := LinkNode{val, nil}
	point.Next = &newNode
}

func (head *LinkNode) Reverse() *LinkNode {
	if head == nil  {
		return head
	}

	var pre *LinkNode
	var next *LinkNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func (link *LinkNode) Print() {
	if link == nil {
		return
	}
	fmt.Println(link.Value)
	for link.Next != nil {
		fmt.Println(link.Next.Value)
		link = link.Next
	}
}

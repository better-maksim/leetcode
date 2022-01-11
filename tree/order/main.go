package order

import (
	"fmt"
	"gitbuh.com/better-maksim/leetcode/tree"
)

func RecurPreOrder(root *tree.Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	RecurPreOrder(root.LeftNode)
	RecurPreOrder(root.RightNode)
}

func RecurInOrder(root *tree.Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	RecurInOrder(root.LeftNode)
	RecurInOrder(root.RightNode)
}

func RecurPosOrder(root *tree.Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	RecurPosOrder(root.LeftNode)
	RecurPosOrder(root.RightNode)
}

func UnRecurPreOrder()  {

}
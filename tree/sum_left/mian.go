package main

import (
	"fmt"
	"gitbuh.com/better-maksim/leetcode/tree"
)

var res = 0

func main() {
	rootNode := tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val: 9,
		},
		Right: &tree.TreeNode{
			Val: 20,
			Left: &tree.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TreeNode{
				Val:   24,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(&rootNode))
}

func sumOfLeftLeaves(root *tree.TreeNode) int{
	orderTree(root, false)
	return res
}

func orderTree(node *tree.TreeNode, isLeft bool) {

	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil && isLeft {
		res += node.Val
	}
	orderTree(node.Left, true)
	orderTree(node.Right, false)
}

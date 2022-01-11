package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rootNode := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   24,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(numColor(&rootNode))
}

func numColor(root *TreeNode) int {
	color := make(map[int]int)
	orderNode(root, &color)
	return len(color)
}

func orderNode(root *TreeNode, color *map[int]int) {
	if root == nil {
		return
	}
	orderNode(root.Left, color)
	(*color)[root.Val] = 1
	orderNode(root.Right, color)

}

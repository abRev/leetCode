package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deep(root *TreeNode, now int) int {
	if root == nil {
		return 0
	}
	now = now*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return now
	}
	return deep(root.Left, now) + deep(root.Right, now)
}

func sumNumbers(root *TreeNode) int {
	now := 0
	return deep(root, now)
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}
	fmt.Println(sumNumbers(root))
}

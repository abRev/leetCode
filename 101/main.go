package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	result := true
	if root == nil {
		return result
	}
	deepSymmetric(root.Left, root.Right, &result)
	return result
}

func deepSymmetric(left, right *TreeNode, result *bool) {
	if *result == false {
		return
	}
	if left == nil && right == nil {
		return
	}
	if (left != nil && right == nil) || (left == nil && right != nil) || left.Val != right.Val {
		*result = false
		return
	}
	deepSymmetric(left.Left, right.Right, result)
	deepSymmetric(left.Right, right.Left, result)
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 3,
	}
	fmt.Println(isSymmetric(root))
}

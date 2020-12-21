package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

}

func insert(root *TreeNode, val int) {
	if root.Val > val {
		if root.Left != nil {
			insert(root.Left, val)
		} else {
			root.Left = &TreeNode{
				Val: val,
			}
		}
	} else {
		if root.Right != nil {
			insert(root.Right, val)
		} else {
			root.Right = &TreeNode{
				Val: val,
			}
		}
	}
}

func copy(root, dest *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		dest.Left = &TreeNode{
			Val: root.Left.Val,
		}
		copy(root.Left, dest.Left)
	}
	if root.Right != nil {
		dest.Right = &TreeNode{
			Val: root.Right.Val,
		}
		copy(root.Right, dest.Right)
	}
}

func main() {

}

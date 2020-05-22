package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := TreeNode{
		Val: preorder[0],
	}
	middle := 0
	for index, value := range inorder {
		if value == preorder[0] {
			middle = index
			break
		}
	}
	head.Left = buildTree(preorder[1:middle+1], inorder[0:middle])
	head.Right = buildTree(preorder[middle+1:], inorder[middle+1:])
	return &head
}

func main() {
	preorder := []int{
		1, 2,
	}
	inorder := []int{
		1, 2,
	}
	head := buildTree(preorder, inorder)
	fmt.Println("------", head)
}

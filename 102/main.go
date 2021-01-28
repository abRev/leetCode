package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	arr := []*TreeNode{}
	if root != nil {
		arr = append(arr, root)
	}
	for len(arr) > 0 {
		vals := []int{}
		temp := []*TreeNode{}
		for _, row := range arr {
			vals = append(vals, row.Val)
			if row.Left != nil {
				temp = append(temp, row.Left)
			}
			if row.Right != nil {
				temp = append(temp, row.Right)
			}
		}
		result = append(result, vals)
		arr = temp
	}
	return result
}

func main() {

}

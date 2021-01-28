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
	status := true
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
		if !status {
			for i := 0; i < len(vals)/2; i++ {
				temp2 := vals[i]
				vals[i] = vals[len(vals)-1-i]
				vals[len(vals)-1-i] = temp2
			}
		}
		result = append(result, vals)
		arr = temp
		status = !status
	}
	return result
}

func main() {

}

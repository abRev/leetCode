/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
        return result
    }
    row := []*TreeNode{root}
    temp := []int{root.Val}
    result = append(result,temp)
    for len(row) > 0 {
        temp := []int{}
        tRow := []*TreeNode{}
        for _,node := range row {
            if node.Left != nil {
                temp = append(temp,node.Left.Val)
                tRow = append(tRow,node.Left)
            }
            if node.Right != nil {
                temp = append(temp,node.Right.Val)
                tRow = append(tRow,node.Right)
            }
        }
        if len(temp) > 0 {
            result = append(result, temp)
        }
        row = tRow
    }
    for i:=0;i<len(result)/2;i++ {
        temp = result[i]
        result[i] = result[len(result)-i-1]
        result[len(result)-i-1] = temp
    }
    return result
}

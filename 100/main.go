/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    result := true
    deep(p,q,&result)
    return result
}

func deep(p, q *TreeNode, result *bool) {
    if *result == false || (p == nil && q==nil){
        return
    }
    if (p == nil && q!= nil) || (q==nil && p !=nil) ||  p.Val != q.Val {
        *result = false
        return 
    }
    deep(p.Left,q.Left,result)
    deep(p.Right,q.Right,result)
}
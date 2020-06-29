/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    deepCal(root, 0)
    return root
}

func deepCal(root *TreeNode,parent int) int {
    // 获取右子树的和
    // 当前值为右子树的和加当前val
    // 左子树为父节点的val+当前节点加当前节点的右子树
    if root == nil {
        return parent
    }
    if root.Right == nil && root.Left == nil{
        root.Val = root.Val + parent
        return root.Val
    }
    right := deepCal(root.Right,parent)
    root.Val = root.Val + right
    return deepCal(root.Left,root.Val)
}

func flatten(root *TreeNode)  {
    deep(root)
}
func deep(root *TreeNode) {
    if root == nil {
        return
    }
    if root.Left != nil {
        deep(root.Left)
        right := root.Right
        left := root.Left
        root.Right = left
        root.Left = nil
        for left.Right != nil {
            left=left.Right
        }
        left.Right = right
    }
    deep(root.Right)
    
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    } else if len(nums) == 1 {
        return &TreeNode{
            Val: nums[0],
        }
    }
    index := len(nums)/2
    root := &TreeNode{
        Val: nums[index],
    }
    root.Left = sortedArrayToBST(nums[:index])
    root.Right = sortedArrayToBST(nums[index+1:])
    return root
}

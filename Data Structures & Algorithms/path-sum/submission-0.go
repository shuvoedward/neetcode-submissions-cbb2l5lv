/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, total int)bool

	dfs = func(root *TreeNode, total int)bool{
		if root == nil{
			return false
		}

		total += root.Val

		if root.Left == nil && root.Right == nil {
			if total == targetSum {
				return true
			}else{
				return false
			}
		}
			
		left := dfs(root.Left, total)
		right := dfs(root.Right, total)

		return left || right
	}

	return dfs(root, 0)
}

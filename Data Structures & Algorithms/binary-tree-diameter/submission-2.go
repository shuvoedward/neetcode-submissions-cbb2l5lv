/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

	var dfs func(root *TreeNode)int
	dfs = func(root *TreeNode)int{
		// diameter from each node
		// goes up the max from left, right

		if root == nil{
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		maxDiameter = max(maxDiameter, l+r)

		return 1 + max(l, r)
	}

	dfs(root)

	return maxDiameter
}

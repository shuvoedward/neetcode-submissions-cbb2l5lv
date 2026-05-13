/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    // max 
	// goodNodes
	
	var dfs func(root *TreeNode, max int)int
	dfs = func(root *TreeNode, max int)int{
		if root == nil{
			return 0
		}

		count := 0
		if max <= root.Val{
			max = root.Val
			count = 1
		}

		return count + dfs(root.Left, max) + dfs(root.Right, max)

	}

	return dfs(root, root.Val)
	
}

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
	goodNodes := 0
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int){
		if root == nil{
			return 
		}

		if max <= root.Val{
			max = root.Val
			goodNodes++
		}

		dfs(root.Left, max)
		dfs(root.Right, max)

	}

	dfs(root, root.Val)
	return goodNodes
}

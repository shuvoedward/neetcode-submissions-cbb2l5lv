/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
   // traverse to leaf node, 
		// leaf when no children
		// if leaf == target, return nil to parent. 
			// parent.left = nil
		// 

		var dfs func(*TreeNode) *TreeNode
		dfs = func(root *TreeNode) *TreeNode{
			if root == nil{
				return nil
			}

			root.Left = dfs(root.Left)
			root.Right = dfs(root.Right)

			if root.Left == nil && root.Right == nil && root.Val == target{
			return nil	
			}

			return root
		}

		return dfs(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
   // val > root, right
   // val < root, left
   // if nil, return val
	
	var dfs func(root *TreeNode, val int)*TreeNode
	dfs = func(root *TreeNode, val int)*TreeNode{
		if root == nil{
			return &TreeNode{
				Val: val,
			}
		}

		if val > root.Val {
			root.Right =  dfs(root.Right, val)
		}else{
			root.Left = dfs(root.Left, val)
		}

		return root
	}
   return dfs(root, val)

}

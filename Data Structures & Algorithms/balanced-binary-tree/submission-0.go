/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Result struct{
	balanced bool
	height int
}
func abs(a int)int{
	if a < 0{
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
    // not height either equal or differ by 1 only.

	var dfs func(root *TreeNode)Result
	// count height, left and right
	// check if it is balanced.
	dfs = func(root *TreeNode)Result{
		if root == nil{
			return Result{true, 0}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		balanced := left.balanced && right.balanced && abs(left.height - right.height) <= 1

		return Result{balanced, 1 + max(left.height, right.height)}
	}

	return dfs(root).balanced 
	

}

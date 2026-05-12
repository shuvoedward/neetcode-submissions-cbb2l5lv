/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {

	res := []int{}

	var recursion func(root *TreeNode)
	
	recursion = func(root *TreeNode){
		if root == nil{
			return 
		}

		recursion(root.Left)
		res = append(res, root.Val)
		recursion(root.Right)
	}

	recursion(root)

	return res

}

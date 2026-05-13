/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
   res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		nodeVal := []int{}
		for range l {
			node := queue[0]
			queue = queue[1:]
			nodeVal = append(nodeVal, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, nodeVal)
		
	}

	return res 
}

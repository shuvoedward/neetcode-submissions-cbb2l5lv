/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
   mp := map[*Node]*Node{}

	cur := head

	for cur != nil {
		mp[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
		cur = cur.Next
	}

	return mp[head] 
}

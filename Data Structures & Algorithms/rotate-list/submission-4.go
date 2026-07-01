/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil{
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0{
		return head
	}

   cur := head
   for range (length - k - 1){
	cur = cur.Next
   }

	newHead := cur.Next
	cur.Next = nil
	tail.Next = head

   return newHead
}

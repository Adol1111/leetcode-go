package solutions

func reverseK(head *ListNode, k int) *ListNode {
	prev := head
	for i := 1; i < k; i++ {
		current := prev.Next
		prev.Next = current.Next
		current.Next = head
		head = current
	}
	return head
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	for node := head; node != nil; node = node.Next {
		count++
	}

	var prevGroupTail *ListNode
	groupHead := head
	for count >= k && groupHead != nil {
		count -= k
		tail := groupHead
		groupHead = reverseK(groupHead, k)
		if prevGroupTail == nil {
			prevGroupTail = tail
			head = groupHead
		} else {
			prevGroupTail.Next = groupHead
			prevGroupTail = tail
		}
		groupHead = tail.Next
	}
	return head
}

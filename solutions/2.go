package solutions

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := l1
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		l1.Val += l2.Val
		if l1.Val >= 10 {
			l1.Val -= 10
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1}
			} else {
				l1.Next.Val++
			}
		}
		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return head
}

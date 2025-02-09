package solutions

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	// 原理，当前节点不参与比较，只比较当前节点的后两项
	// 如果当前节点的后两项值不同，则cur往后移动一个节点 (不是2个，第2个节点还没有和第3个节点比较)
	// 如果当前节点的后两项值相同，则让cur.next移动到下一个不同值的节点
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 如果出现重复则跳过，找下一个不同值的节点
func DeleteDuplicates1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head
	for fast != nil {
		if (slow == dummy || slow.Val != fast.Val) && (fast.Next == nil || fast.Next.Val != fast.Val) {
			slow.Next = fast
			slow = fast
		}
		next := fast.Next
		for next != nil && next.Val == fast.Val {
			next = next.Next
		}
		fast = next
	}
	if slow == dummy {
		return nil
	} else {
		slow.Next = nil
	}
	return dummy.Next
}

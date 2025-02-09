package solutions

func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	fast := dummy
	for fast.Next != nil {
		if fast.Next.Val >= x {
			break
		}
		fast = fast.Next
	}
	slow := fast
	prev := slow
	for {
		fast = prev.Next
		if fast == nil {
			break
		}
		if fast.Val < x {
			prev.Next = fast.Next
			fast.Next = slow.Next
			slow.Next = fast
			slow = slow.Next
		}
		prev = fast
	}
	return dummy.Next
}

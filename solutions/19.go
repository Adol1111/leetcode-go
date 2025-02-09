package solutions

// 使用栈来实现
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var stack []*ListNode
	dummy := &ListNode{Next: head}

	for node := dummy; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	node := stack[len(stack)-1-n]
	node.Next = node.Next.Next
	return dummy.Next
}

// 计算链表长度
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	count := 1
	node := head
	for node.Next != nil {
		node = node.Next
		count++
	}
	if count == n {
		return head.Next
	}
	node = head
	for i := 1; i < count-n; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}

// 快慢指针, 利用快慢指针的差 = n，当快指针到达尾部时，慢指针刚好到达倒数第n个节点
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

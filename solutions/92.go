package solutions

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	node := head
	for i := 1; i < left; i++ {
		prev = node
		node = node.Next
	}

	node = reverseK(node, right-left+1)
	if prev != nil {
		prev.Next = node
	} else {
		head = node
	}

	return head
}

// a->b->c->d->e
// 2, 4
// 2 = b
// 4 = d
// leftNode = a
// 1. prev = a, node = b

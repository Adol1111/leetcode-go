package solutions

// 两个有序链表进行合并
func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			head.Next = left
			left = left.Next
		} else {
			head.Next = right
			right = right.Next
		}
		head = head.Next
	}
	if left != nil {
		head.Next = left
	}
	if right != nil {
		head.Next = right
	}
	return dummy.Next
}

// 归并排序，每次对半拆开，拆成左右两个链表，再把左右两个链表再拆开，直到不可拆为止
// 然后把左右两个链表各自排序（递归调用），最后把两个有序链表进行合并
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil // 断开右侧节点

	left := SortList(head)
	right := SortList(mid)
	return merge(left, right)
}

// 插入排序，时间复杂度不满足要求
func SortList1(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	last := dummy.Next
	for last != nil && last.Next != nil {
		cur := last.Next
		if last.Val < cur.Val {
			last = cur
			continue
		}
		prev := dummy
		node := dummy.Next
		for node != cur {
			if node.Val >= cur.Val {
				last.Next = cur.Next
				cur.Next = node
				prev.Next = cur
				break
			}
			node = node.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}

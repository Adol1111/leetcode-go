package solutions

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		head = list1
		list1 = list1.Next
	}

	merged := head
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil {
			merged.Next = list2
			break
		}
		if list2 == nil {
			merged.Next = list1
			break
		}
		if list1.Val > list2.Val {
			merged.Next = list2
			list2 = list2.Next
		} else {
			merged.Next = list1
			list1 = list1.Next
		}
		merged = merged.Next
	}
	return head
}

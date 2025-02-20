package solutions

func MergeKList(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	left := MergeKList(lists[:n/2])
	right := MergeKList(lists[n/2:])
	return merge(left, right)
}

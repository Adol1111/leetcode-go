package solutions

// 核心思想为快慢指针，快指针先走k步，然后快慢指针同时走，快指针到达尾部时，慢指针刚好到达倒数第k个节点
// 这里取余数，因为k如果超过链表长度后，只要再走一圈即可，多余的圈数不会影响结果
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	count := 0
	for i := 0; i < k; i++ {
		count++
		fast = fast.Next
		if fast == nil {
			// 修正k值为1圈+余数, 因为此时i等于一圈，除非重置i为0
			k = k%count + count
			fast = head
		}
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}

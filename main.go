package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// result := solutions.Trap([]int{4, 2, 0, 3, 2, 5})

	l1 := &solutions.ListNode{
		Val: 1,
		Next: &solutions.ListNode{
			Val: 2,
			Next: &solutions.ListNode{
				Val: 3,
				Next: &solutions.ListNode{
					Val: 4,
					Next: &solutions.ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	result := solutions.ReverseKGroup(l1, 2)
	fmt.Printf("result=%v", result)
}

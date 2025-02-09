package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// result := solutions.Trap([]int{4, 2, 0, 3, 2, 5})

	l1 := &solutions.TreeNode{
		Val: 2,
		Left: &solutions.TreeNode{
			Val: -1,
		},
	}
	result := solutions.MaxPathSum(l1)
	fmt.Printf("result=%v", result)
}

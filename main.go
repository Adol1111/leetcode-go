package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// result := solutions.Trap([]int{4, 2, 0, 3, 2, 5})

	// l1 := &solutions.TreeNode{
	// 	Val: 2,
	// 	Left: &solutions.TreeNode{
	// 		Val: -1,
	// 	},
	// }
	result := solutions.FindMin([]int{3, 4, 5, 1, 2})
	fmt.Printf("result=%v", result)
}

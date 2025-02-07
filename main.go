package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// result := solutions.Trap([]int{4, 2, 0, 3, 2, 5})
	result := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	solutions.Rotate(result)
	fmt.Printf("result=%v", result)
}

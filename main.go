package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// result := solutions.Trap([]int{4, 2, 0, 3, 2, 5})
	result := solutions.Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	fmt.Printf("result=%v", result)
}

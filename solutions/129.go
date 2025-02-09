package solutions

func getNumbers(root *TreeNode, number int) []int {
	if root == nil {
		return []int{}
	}
	number = number*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return []int{number}
	}
	var result []int
	result = append(result, getNumbers(root.Left, number)...)
	result = append(result, getNumbers(root.Right, number)...)
	return result
}

func SumNumbers(root *TreeNode) int {
	nums := getNumbers(root, 0)
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

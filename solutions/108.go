package solutions

func SortedArrayToBST(nums []int) *TreeNode {
	var help func(nums []int, left, right int) *TreeNode
	help = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = help(nums, left, mid-1)
		root.Right = help(nums, mid+1, right)
		return root
	}
	return help(nums, 0, len(nums)-1)
}

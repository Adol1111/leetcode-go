package solutions

func TwoSum2(nums []int, target int) []int {
	ans := make([]int, 2)
	hashTable := make(map[int]int)

	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			ans[0] = j
			ans[1] = i
			return ans
		}
		hashTable[num] = i
	}
	return nil
}

package solutions

func TwoSum2(nums []int, target int) []int {
	ans := make([]int, 2)
	hashTable := make(map[int]int)

	for i, num := range nums {
		// 如果减去当前的值在map中存在，那么就找到了，返回结果
		if j, ok := hashTable[target-num]; ok {
			ans[0] = j
			ans[1] = i
			return ans
		}
		// 把下标存入map
		hashTable[num] = i
	}
	return nil
}

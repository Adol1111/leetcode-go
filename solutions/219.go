package solutions

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if pos, ok := m[num]; ok && i-pos <= k {
			return true
		}
		m[num] = i
	}
	return false
}

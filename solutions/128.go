package solutions

func LongestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	longestConsecutive := 0
	for num := range m {
		if !m[num-1] {
			current := num
			currentConsecutive := 1
			for m[current+1] {
				current++
				currentConsecutive++
			}
			longestConsecutive = max(longestConsecutive, currentConsecutive)
		}
	}
	return longestConsecutive
}

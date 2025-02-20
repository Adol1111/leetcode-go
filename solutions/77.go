package solutions

func Combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			// 拷贝结果，path需要递归使用，所以需要拷贝
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			// 回溯
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return res
}

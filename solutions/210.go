package solutions

func FindOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		// 状态表，0表示未上课，1表示进行中，2表示已上课完成, 未出现循环依赖
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
	)

	// 深度优先，找到一门课就把所有依赖都找出来，并放到结果中
	var dfs func(int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			return false
		}
		if visited[node] == 2 {
			return true
		}
		visited[node] = 1

		for _, nextNode := range edges[node] {
			if !dfs(nextNode) {
				return false
			}
		}
		// 所有依赖的课都已完成，标记为已完成
		visited[node] = 2
		result = append(result, node)
		return true
	}

	for _, info := range prerequisites {
		// 收集依赖关系, 课程[0]依赖课程[1]
		edges[info[0]] = append(edges[info[0]], info[1])
	}

	for i := 0; i < numCourses && valid; i++ {
		// 如果不是0 则说明已收集过，不需要重复处理
		if visited[i] == 0 && !dfs(i) {
			return []int{}
		}
	}

	return result
}

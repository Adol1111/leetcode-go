package solutions

// 等价于判断是否存在循环图
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 构建邻接表
	adjacency := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		// 表示 课程[0] 依赖 课程[1]
		adjacency[prerequisite[0]] = append(adjacency[prerequisite[0]], prerequisite[1])
	}
	// 记录节点状态
	// 0 未访问
	// 1 访问中
	// 2 已访问
	status := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(node int) bool {
		// 如果还在访问中，说明出现了循环访问，返回 false
		if status[node] == 1 {
			return false
		}
		if status[node] == 2 {
			return true
		}
		status[node] = 1
		for _, nextNode := range adjacency[node] {
			// 如果下一个节点已经访问过了，说明出现了循环访问，返回 false
			if !dfs(nextNode) {
				return false
			}
		}
		// 所有依赖的节点都已遍历，则设置为已访问
		status[node] = 2
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

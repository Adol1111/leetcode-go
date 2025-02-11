package solutions

func FindOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
	)

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = 1

		for _, nextNode := range edges[node] {
			if visited[nextNode] == 0 { // 未上课，继续
				dfs(nextNode)
				if !valid {
					return
				}
			} else if visited[nextNode] == 1 { // 进行中，说明有环，返回 false
				valid = false
				return
			}
		}
		visited[node] = 2 // 无任何依赖，可以完成，加入结果集
		result = append(result, node)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

package solutions

func isAdjacent(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff == 1
}

func MinMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	adjacency := make([][]int, len(bank))
	endIndex := -1
	for i := 0; i < len(bank); i++ {
		if bank[i] == endGene {
			endIndex = i
		}

		for j := i + 1; j < len(bank); j++ {
			if isAdjacent(bank[i], bank[j]) {
				adjacency[i] = append(adjacency[i], j)
				adjacency[j] = append(adjacency[j], i)
			}
		}
	}
	if endIndex == -1 {
		return -1
	}

	visited := make([]bool, len(bank))
	queue := []int{}
	for i, s := range bank {
		if isAdjacent(startGene, s) {
			queue = append(queue, i)
			visited[i] = true
		}
	}

	for step := 1; len(queue) > 0; step++ {
		tmp := queue
		queue = nil // 清空queue，重新补充
		for _, cur := range tmp {
			if cur == endIndex {
				return step
			}
			// 获取当前节点的下一个节点，遍历所有下一个节点，如果没有访问过，则加入队列中
			for _, next := range adjacency[cur] {
				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}

		}
	}
	return -1
}

// 广度优先算法
func MinMutation2(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}

	if !bankSet[endGene] {
		return -1
	}

	queue := []string{startGene}

	for step := 0; len(queue) > 0; step++ {
		tmp := queue              // 防止读到新的queue
		for _, cur := range tmp { // 遍历queue中所有的可能性后，再把step+1
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x { // 每次只改一位，如果当前位和目标位相同，就不修改
						// 新生成的基因如果==目标基因，就返回步数
						// 否则如果在bank中，就加入队列中，不再则忽略
						next := cur[:i] + string(y) + cur[i+1:]
						if next == endGene {
							return step + 1
						}
						if bankSet[next] {
							queue = append(queue, next)
							delete(bankSet, next)
						}
					}
				}
			}
		}
	}
	return -1
}

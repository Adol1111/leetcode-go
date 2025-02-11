package solutions

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// DFS 深度优先，每次都是clone完整节点，如果发现有Neighbors，则利用递归去clone Neighbors节点
func CloneGraph(node *GraphNode) *GraphNode {
	visited := make(map[*GraphNode]*GraphNode)

	var dfsClone func(node *GraphNode) *GraphNode
	dfsClone = func(node *GraphNode) *GraphNode {
		if node == nil {
			return nil
		}
		if node, ok := visited[node]; ok {
			return node
		}
		clone := &GraphNode{Val: node.Val}
		visited[node] = clone
		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfsClone(neighbor))
		}
		return clone
	}
	return dfsClone(node)
}

// BFS 广度优先，先创建节点, 再处理Neighbors
// 如果发现visited中没有，则创建节点并加入队列中,
// 如果发现visited中有，那么加入已存在节点的Neighbors中
func CloneGraph2(node *GraphNode) *GraphNode {
	if node == nil {
		return node
	}
	visited := map[*GraphNode]*GraphNode{}

	queue := []*GraphNode{node}
	visited[node] = &GraphNode{Val: node.Val, Neighbors: make([]*GraphNode, 0)}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, neighbor := range cur.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &GraphNode{Val: neighbor.Val, Neighbors: make([]*GraphNode, 0)}
				queue = append(queue, neighbor)
			}
			// 当前节点一定在visited中，因为当前节点是从queue中取出的首个节点
			// 如果节点没有在visited中，那么会先加入到visited中再加入到queue中
			// 既然queue中存在，那visited中一定存在
			visited[cur].Neighbors = append(visited[cur].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}

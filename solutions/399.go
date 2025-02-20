package solutions

// union find 是求最小生成树的一种算法
type unionFind struct {
	parent []int
	weight []float64
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	weight := make([]float64, n)
	for i := 0; i < n; i++ {
		// 初始条件
		// 1. id相同的节点，parent都是自己
		// 2. 自己等于自己，所以权重是1
		parent[i] = i
		weight[i] = 1.0
	}
	return &unionFind{
		parent: parent,
		weight: weight,
	}
}

func (uf *unionFind) find(x int) int {
	// 如果x的parent已经是自己了，说明已经建立连接直接返回即可
	// 如果没有建立连接，递归调用find函数，建立连接
	// 此时把x的parent设置为根节点，并把自身权重*父节点的权重，循环几次后就得到了自身到根节点的权重
	if uf.parent[x] != x {
		origin := uf.parent[x]
		uf.parent[x] = uf.find(uf.parent[x])
		uf.weight[x] *= uf.weight[origin]
	}
	return uf.parent[x]
}

func (uf *unionFind) getConnectedResult(x int, y int) float64 {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return uf.weight[x] / uf.weight[y]
	}
	return -1.0
}

func (uf *unionFind) union(x int, y int, value float64) {
	// 都连接到根节点，比如 a/b = 1 , a/c = 2 => a=b, a=2c, 根节点就是a
	rootX := uf.find(x)
	rootY := uf.find(y)

	// 如果存在相同跟节点说明已经建立了连接，直接返回
	// 开始时，a和b无关联，此时 rootX 和 rootY 都是自己
	if rootX == rootY {
		return
	}

	// 比如 a=2b, c=3a => c=6b
	// 1. 开始时，a和b无关联，此时权重为 value (此时rootX=a, rootY=b, parent['a']='b', weight['a']=2, weight['b']保持不变)
	// 2. 需要建立a和c的连接, 此时 (x=c, y=a, rootX=c, rootY=b, parent['c']='b')
	//    此时权重为 weight['c'] = 3 * weight['a'] / weight['c'] = 6
	uf.parent[rootX] = rootY
	uf.weight[rootX] = value * uf.weight[y] / uf.weight[x]
}

// 使用union find算法
func CalcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	// 极端情况 所有equations节点都不关联，此时UnionFind的长度就是equations的两倍
	unionFind := newUnionFind(len(equations) * 2)
	m := make(map[string]int, len(equations)*2)

	id := 0
	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		a := equation[0]
		b := equation[1]

		// 给节点创建唯一的id编号，并将节点缓存到map中，如果遇到相同的节点，保持id一致
		if _, ok := m[a]; !ok {
			m[a] = id
			id++
		}

		if _, ok := m[b]; !ok {
			m[b] = id
			id++
		}

		// 给两个节点id 建立连接
		unionFind.union(m[a], m[b], values[i])
	}

	var result []float64
	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]
		ida, ok1 := m[a]
		idb, ok2 := m[b]
		if !ok1 || !ok2 {
			result = append(result, -1.0)
		} else {
			result = append(result, unionFind.getConnectedResult(ida, idb))
		}
	}
	return result
}

// 使用广度优先算法
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := map[string]int{}
	id := 0
	// 把节点名称转为id
	for _, equation := range equations {
		a, b := equation[0], equation[1]
		if _, ok := m[a]; !ok {
			m[a] = id
			id++
		}
		if _, ok := m[b]; !ok {
			m[b] = id
			id++
		}
	}

	// 构建图
	type edge struct {
		to     int
		weight float64
	}

	// 这里使用id做节点，而不是节点名称
	graph := make([][]edge, id)
	for _, equation := range equations {
		a, b := equation[0], equation[1]
		// 构建双向图，每个节点都有一个到另一个节点的边 a->b 和 b->a，两者的value是倒数关系
		graph[m[a]] = append(graph[m[a]], edge{m[b], values[0]})
		graph[m[b]] = append(graph[m[b]], edge{m[a], 1 / values[0]})
	}

	// 广度遍历，从start开始，找到end节点，返回从start到end的权重，即所有经过的边的权重相乘
	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == end {
				return ratios[v]
			}
			for _, e := range graph[v] {
				// 如果已经计算过到start的权重值，则忽略, 否则收集到queue中
				if w := e.to; ratios[w] == 0 {
					// ratios[w] 表示从start到w的权重 = 从start到v的权重 * v到w的权重
					ratios[w] = ratios[v] * e.weight
					queue = append(queue, w)
				}
			}
		}
		return -1
	}

	var result []float64
	for _, query := range queries {
		a, b := query[0], query[1]
		ida, ok1 := m[a]
		idb, ok2 := m[b]
		// 如果任意节点此前没有出现过，说明无法计算
		if !ok1 || !ok2 {
			result = append(result, -1.0)
		} else {
			result = append(result, bfs(ida, idb))
		}
	}
	return result
}

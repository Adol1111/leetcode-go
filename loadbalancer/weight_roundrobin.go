package loadbalancer

type WeightRoundRobinPicker struct {
	servers       []*Server
	index         int
	currentWeight int
	maxWeight     int
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func gcdServers(servers []*Server) int {
	size := len(servers)
	if size == 0 {
		return 0
	}
	gcd := servers[0].Weight
	for i := 1; i < size; i++ {
		gcd = Gcd(gcd, servers[i].Weight)
	}
	return gcd
}

func maxWeight(servers []*Server) int {
	max := 0
	for _, server := range servers {
		if server.Weight > max {
			max = server.Weight
		}
	}
	return max
}

func (p *WeightRoundRobinPicker) Next() *Server {
	if len(p.servers) == 0 {
		return nil
	}
	// 原理
	// 根据最大公约数原理，完整一轮会经过小轮的次数为 最大权重 / gcd 次（因为是最大公约数，所以次数是整数）
	// 当前权重 = 最大权重 - N * gcd，当执行次数 N >= (最大权重-wi) / gcd 次时,
	// 当前权重 <= 最大权重 - (最大权重-wi) / gcd * gcd = wi ，此时si一定会被选中,
	// 因此被选中的次数为 总次数 - 被跳过次数 = (最大权重/gcd) - (最大权重-wi)/gcd = wi/gcd 次
	//
	// 综上所述，我们在index回到0时算作一小轮结束，重新计算当前权重值，当前权重 -= gcd, 如果当前权重 <= 0，则重置为最大权重
	// 每一小轮都会遍历所有节点，如果当前节点的权重大于该节点的权重，则当前节点被选中, 否则则跳过该节点取下一个节点
	for {
		p.index = (p.index + 1) % len(p.servers)
		if p.index == 0 {
			p.currentWeight -= gcdServers(p.servers)
			if p.currentWeight <= 0 {
				p.currentWeight = p.maxWeight
			}
		}
		if p.servers[p.index].Weight >= p.currentWeight {
			break
		}
	}
	return p.servers[p.index]
}

func (p *WeightRoundRobinPicker) Name() string {
	return "weight_round_robin"
}

func NewWeightRoundRobinPicker(servers []*Server) Picker {
	return &WeightRoundRobinPicker{
		servers:       servers,
		index:         -1,
		currentWeight: 0,
		maxWeight:     maxWeight(servers),
	}
}

func init() {
	RegisterPickerProvider(WeightRoundRobinPickerType, PickerProviderFunc(NewWeightRoundRobinPicker))
}

package loadbalancer

type SmoothWeightRoundRobinPicker struct {
	servers []*Server
	index   int
}

func (p *SmoothWeightRoundRobinPicker) Next() *Server {
	// 原理
	// 每次所有节点的当前权重都增加了 server.weight，取当前权重最大的节点作为选中节点，如果并列最大取第一个
	// 节点选中后，当前权重减去 total weight
	// 一轮完成时，一个节点会增加 权重总和次 * 节点权重，而减去的权重为被选中的次数 * 权重总和 两者是相等的
	// 也就是说被选中的次数=当前节点的权重，当一轮结束时，所有节点的当前权重都会回到0
	// 由于每次被选中后，权重会减少total weight，所以保证了下次选中节点的概率被降低，从而实现平滑的负载均衡
	var pickedServer *Server
	for _, server := range p.servers {
		server.currentWeight += server.Weight
		if pickedServer == nil || server.currentWeight > pickedServer.currentWeight {
			pickedServer = server
		}
	}

	if pickedServer != nil {
		for _, server := range p.servers {
			pickedServer.currentWeight -= server.Weight
		}
	}
	return pickedServer
}

func (p *SmoothWeightRoundRobinPicker) Name() string {
	return "smooth_weight_round_robin"
}

func NewSmoothWeightRoundRobinPicker(servers []*Server) Picker {
	return &SmoothWeightRoundRobinPicker{
		servers: servers,
		index:   0,
	}
}

func init() {
	RegisterPickerProvider(SmoothWeightRoundRobinPickerType, PickerProviderFunc(NewSmoothWeightRoundRobinPicker))
}

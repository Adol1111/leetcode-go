package loadbalancer

type SmoothWeightRoundRobinPicker struct {
	servers []*Server
	index   int
}

func NewSmoothWeightRoundRobinPicker(servers []*Server) Picker {
	return &SmoothWeightRoundRobinPicker{
		servers: servers,
		index:   0,
	}
}

func (p *SmoothWeightRoundRobinPicker) Next() *Server {
	var pickedServer *Server
	for _, server := range p.servers {
		server.currentWeight += server.Weight
		if pickedServer == nil || server.currentWeight > pickedServer.currentWeight {
			pickedServer = server
		}
	}

	if pickedServer != nil {
		totalWeight := 0
		for _, server := range p.servers {
			totalWeight += server.Weight
		}
		pickedServer.currentWeight -= totalWeight
	}
	return pickedServer
}

func (p *SmoothWeightRoundRobinPicker) Name() string {
	return "smooth_weight_round_robin"
}

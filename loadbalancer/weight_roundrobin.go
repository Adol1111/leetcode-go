package loadbalancer

type Picker interface {
	Next() *Server
	Name() string
}

type Server struct {
	IP            string
	Weight        int
	currentWeight int
}

type WeightRoundRobinPicker struct {
	servers       []*Server
	index         int
	currentWeight int
}

func NewWeightRoundRobinPicker(servers []*Server) Picker {
	return &WeightRoundRobinPicker{
		servers:       servers,
		index:         -1,
		currentWeight: 0,
	}
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
	for {
		p.index = (p.index + 1) % len(p.servers)
		if p.index == 0 {
			p.currentWeight -= gcdServers(p.servers)
			if p.currentWeight <= 0 {
				p.currentWeight = maxWeight(p.servers)
				if p.currentWeight == 0 {
					return nil
				}
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

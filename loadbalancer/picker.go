package loadbalancer

type PickerType string

const (
	WeightRoundRobinPickerType       PickerType = "weight_round_robin"
	SmoothWeightRoundRobinPickerType PickerType = "smooth_weight_round_robin"
)

var (
	pickerProviderRegistry map[PickerType]PickerProvider = make(map[PickerType]PickerProvider)
)

type Server struct {
	IP            string
	Weight        int
	currentWeight int
}

type Picker interface {
	Next() *Server
}

type PickerProvider interface {
	NewPicker(servers []*Server) Picker
}

type PickerProviderFunc func(servers []*Server) Picker

func (f PickerProviderFunc) NewPicker(servers []*Server) Picker {
	return f(servers)
}

func RegisterPickerProvider(t PickerType, provider PickerProvider) {
	pickerProviderRegistry[t] = provider
}

func NewPicker(t PickerType, servers []*Server) Picker {
	pickerProvider, ok := pickerProviderRegistry[t]
	if !ok {
		return nil
	}
	return pickerProvider.NewPicker(servers)
}

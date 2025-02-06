package loadbalancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmoothWeightRoundRobin(t *testing.T) {
	servers := []*Server{
		{
			IP:     "a",
			Weight: 1,
		},
		{
			IP:     "b",
			Weight: 2,
		},
		{
			IP:     "c",
			Weight: 5,
		},
	}

	total := 0
	for _, server := range servers {
		total += server.Weight
	}

	picker := NewSmoothWeightRoundRobinPicker(servers)

	var result []string
	for i := 0; i < total; i++ {
		result = append(result, picker.Next().IP)
	}

	// a b c
	// 0 0 0
	// 1 2 -3(5-8) // c
	// 2 -4(4-8) 2 // b
	// 3 -2 -1(7-8) // c
	// -4(4-8) 0 4 // a
	// -3 2 1(9-8) // c
	// -2 4 -2(6-8) // c
	// -1 -2(6-8) 3 // b
	// 0 0 0(8-8) // c
	assert.Equal(t, []string{
		"c", "b", "c", "a", "c", "c", "b", "c",
	}, result)
}

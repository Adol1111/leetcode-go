package loadbalancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightRoundRobin(t *testing.T) {
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

	var result []string
	picker := NewWeightRoundRobinPicker(servers)
	for i := 0; i < 8; i++ {
		result = append(result, picker.Next().IP)
	}

	// a b c (current, index, next_index)
	// 1 2 5(5, 2, 0) // c
	// 1 2 5(4, 2, 0) // c
	// 1 2 5(3, 2, 0) // c
	// 1 2 5(2, 1, 2) // b
	// 1 2 5(2, 2, 0) // c
	// 1 2 5(1, 0, 1) // a
	// 1 2 5(1, 1, 2) // b
	// 1 2 5(1, 2, 0) // c
	assert.Equal(t, result, []string{"c", "c", "c", "b", "c", "a", "b", "c"})
}

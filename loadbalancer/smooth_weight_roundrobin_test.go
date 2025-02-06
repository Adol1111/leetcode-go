package loadbalancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmoothWeightRoundRobin(t *testing.T) {
	cases := []struct {
		name    string
		servers []*Server
		expect  []string
	}{
		{
			name:    "empty",
			servers: []*Server{},
			expect:  []string(nil),
		},
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
		{
			name: "a=1, b=2, c=5",
			servers: []*Server{
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
			},
			expect: []string{"c", "b", "c", "a", "c", "c", "b", "c"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			total := 0
			for _, server := range c.servers {
				total += server.Weight
			}
			// at least run once when total is 0
			if total == 0 {
				total = len(c.servers)

				if total == 0 {
					total = 1
				}
			}

			var result []string
			picker := NewSmoothWeightRoundRobinPicker(c.servers)
			for i := 0; i < total; i++ {
				picked := picker.Next()
				if picked != nil {
					result = append(result, picked.IP)
				}
			}

			assert.Equal(t, c.expect, result)
		})
	}
}

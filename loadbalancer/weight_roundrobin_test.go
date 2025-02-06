package loadbalancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightRoundRobin(t *testing.T) {
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
		// a(1) b(2) c(5), gcd = 1
		// curent weight 5, pick c
		// curent weight 4, pick c
		// curent weight 3, pick c
		// curent weight 2, pick b, c
		// curent weight 1, pick a, b, c
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
			expect: []string{"c", "c", "c", "b", "c", "a", "b", "c"},
		},

		// a(2) b(2) c(4), gcd = 2
		// curent weight 4, pick c
		// curent weight 2, pick a, b, c
		// curent weight 0 -> 4, pick c
		// curent weight 2, pick a, b, c
		{
			name: "a=2, b=2, c=4",
			servers: []*Server{
				{
					IP:     "a",
					Weight: 2,
				},
				{
					IP:     "b",
					Weight: 2,
				},
				{
					IP:     "c",
					Weight: 4,
				},
			},
			expect: []string{"c", "a", "b", "c", "c", "a", "b", "c"},
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
			picker := NewWeightRoundRobinPicker(c.servers)
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

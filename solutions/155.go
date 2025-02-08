package solutions

import "math"

type MinStack struct {
	data []int
	min  []int
}

func MinStackConstructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  []int{math.MaxInt32},
	}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	top := s.min[len(s.min)-1]
	s.min = append(s.min, min(top, val))
}

func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.min) == 0 {
		return 0
	}
	return s.min[len(s.min)-1]
}

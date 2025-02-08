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

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	top := this.min[len(this.min)-1]
	this.min = append(this.min, min(top, val))
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

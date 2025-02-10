package solutions

import "math/rand"

type RandomizedSet struct {
	data    []int
	indices map[int]int
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{
		data:    []int{},
		indices: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indices[val]
	if !ok {
		return false
	}
	last := len(this.data) - 1
	this.data[index] = this.data[last]
	this.indices[this.data[last]] = index
	this.data = this.data[:last]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

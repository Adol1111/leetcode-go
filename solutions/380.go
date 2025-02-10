package solutions

import "math/rand"

type RandomizedSet struct {
	data    []int
	indices map[int]int // 存放元素在 data 中的索引
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
	this.data[index] = this.data[last]    // 将最后一个元素移动到 index 位置
	this.indices[this.data[last]] = index // 更新最后一个元素的索引
	this.data = this.data[:last]          // 删除最后一个元素
	delete(this.indices, val)             // 删除 val 的索引
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

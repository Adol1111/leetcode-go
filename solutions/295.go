package solutions

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	queMin, queMax Heap
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	// 原理，让queue分成两半，比中位数小的放到queueMin， 比中位数大的放到queueMax
	// 即 queue = queueMin + queueMax
	// 首先这两个堆堆顶都是最小值
	// queueMin存的是比中位数小的值，为了让最大值存放在堆顶，所以需要取负值
	// queueMax存的是比中位数大的值，堆顶是最小值，所以正好是中位数
	// 每当插入新的数时，需要不停移动中位数，要从queueMin->queueMax或者queueMax->queueMin

	// 比如 min堆中是 1 max堆中是5，现在插入-2，结果如下
	// -1 2, 5
	// 比如 min堆中是 1 max堆中是5，现在插入3，结果如下
	// -1, 3 5 -> -3 -1 -> 5
	// 比如 min堆中是 1 3 max堆中是5，现在要插入2，结果如下
	// -3 -1 -2 , 5 -> -2 -1, 3 5
	// 比如 min堆中是 1 3 max堆中是5，现在要插入2，结果如下
	// -3 -1 , 4 5
	minQ, maxQ := &mf.queMin, &mf.queMax
	// num比最小堆的堆顶（即中位数）小，则放入min堆堆顶
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			// 由于最小堆多了一个数，需要把最小堆中的堆顶移动到最大堆中，保证两个堆的平衡
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		// 同理，把比中位数大的值插入最大堆中，然后把最大堆中的堆顶，移动到最小堆中，保持两个堆的平衡
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	// 奇数情况，中位数是最小堆的堆顶
	// 偶数情况，中位数是(最大堆+最小堆的堆顶)/2，因为最小堆是取的负值，所以用-
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type Heap struct{ sort.IntSlice }

func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

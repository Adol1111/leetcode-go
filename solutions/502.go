package solutions

import (
	"container/heap"
	"sort"
)

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct {
		profit  int
		capital int
	}
	arr := make([]pair, n)
	for i := range profits {
		arr[i] = pair{
			capital: capital[i],
			profit:  profits[i],
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].capital < arr[j].capital
	})

	h := &maxHeap{}
	for cur := 0; k > 0; k-- {
		// 把所有成本小于w的项目的利润都放到堆中，然后从中取最大值大的利润（即弹出堆顶元素）
		// 由于arr已经按照成本递增排列，所以只需要判断当前项目的成本是否小于w即可
		// 如果cur的成本>w，则暂停加入堆中，直到w>=arr[cur].capital
		for cur < n && arr[cur].capital <= w {
			heap.Push(h, arr[cur].profit)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

// 实现了go内置的堆接口
type maxHeap struct{ sort.IntSlice }

func (h maxHeap) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *maxHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *maxHeap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

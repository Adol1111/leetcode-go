package solutions

import "container/heap"

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)

	hp := &minHeap{
		data:  make([]pair, 0),
		nums1: nums1,
		nums2: nums2,
	}

	// 初始化堆, 先把较小的结果保存在堆中，然后每弹出一个再补充一个，这样速度会比较快
	// 由于已排序，{i,j+1}一定比{i,j}大，且比{i,j+2}小，因此补充的数据，只需要补充 {i,j+1}即可
	// 不必排列所有数据
	for i := 0; i < k && i < m; i++ {
		hp.data = append(hp.data, pair{i, 0})
	}

	var result [][]int

	for hp.Len() > 0 && len(result) < k {
		// 每弹出一个最小值，再补充一个值，{i,j+1} 一定大于 {i,j}
		p := heap.Pop(hp).(pair)
		i, j := p.i, p.j
		result = append(result, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(hp, pair{i, j + 1})
		}
	}

	return result
}

type pair struct{ i, j int }

type minHeap struct {
	data         []pair
	nums1, nums2 []int
}

func (h minHeap) Len() int { return len(h.data) }
func (h minHeap) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h minHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minHeap) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *minHeap) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

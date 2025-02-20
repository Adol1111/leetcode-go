package solutions

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	// quickSelect 是从小到大排列，第K大的数就是第N-K+1小的数
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[k]
	}
	// 随机选择一个数作为partition, 这里取最左边和随机效果一样
	partition := nums[left]
	i := left - 1
	j := right + 1
	for i < j {
		// 把比partition小的数放到partition的左边，大的数放到partition的右边
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 此时 j 表示partition的位置，此时从小到大排列，k表示第K小的数，由于前面传入的是N-K所以也是第K大的数
	// 如果K在j的左边，则在j的左边继续查找，否则在j的右边继续查找
	if k <= j {
		return quickSelect(nums, left, j, k)
	} else {
		return quickSelect(nums, j+1, right, k)
	}
}

func FindKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildHeap(nums, heapSize)

	// 求第K个，即删除K-1个堆顶元素，每次删除后让堆平衡
	for i := 0; i < k-1; i++ {
		// 删除堆顶元素，即把堆顶和堆底交换，并减少堆大小，然后让堆保持稳定
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	// 由于删除了K-1次，因此堆顶即为第K个最大数
	return nums[0]
}

func buildHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	// i的左右子节点，当前最大数是堆的父节点，即i本身
	left, right, largest := i*2+1, i*2+2, i
	// 3个数（左右父）比较大小，找到最大的序号
	// 如果左子节点比父节点大，则largest为左子节点
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	// 如果右子节点比父节点大，则largest为右子节点
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		// 从堆顶向下递归，直到堆稳定
		maxHeapify(nums, largest, heapSize)
	}
}

package solutions

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*node
	head     *node
	tail     *node
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
	}
}

func (this *LRUCache) moveToHead(node *node) {
	if node == this.head {
		return
	}

	if node == this.tail {
		this.tail = node.prev
		this.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.next = this.head
	node.prev = nil
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) addToHead(node *node) {
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
}

func (this *LRUCache) removeTail() *node {
	node := this.tail
	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}
	return node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &node{
			key:   key,
			value: value,
		}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

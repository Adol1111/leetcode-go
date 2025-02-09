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

func (c *LRUCache) moveToHead(node *node) {
	if node == c.head {
		return
	}

	if node == c.tail {
		c.tail = node.prev
		c.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.next = c.head
	node.prev = nil
	c.head.prev = node
	c.head = node
}

func (c *LRUCache) addToHead(node *node) {
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		node.next = c.head
		c.head.prev = node
		c.head = node
	}
}

func (c *LRUCache) removeTail() *node {
	node := c.tail
	c.tail = c.tail.prev
	if c.tail == nil {
		c.head = nil
	} else {
		c.tail.next = nil
	}
	return node
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; !ok {
		node := &node{
			key:   key,
			value: value,
		}
		c.cache[key] = node
		c.addToHead(node)
		c.size++
		if c.size > c.capacity {
			removed := c.removeTail()
			delete(c.cache, removed.key)
			c.size--
		}
	} else {
		node := c.cache[key]
		node.value = value
		c.moveToHead(node)
	}
}

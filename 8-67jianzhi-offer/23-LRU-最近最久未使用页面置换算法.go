package jianzhi

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkNode
	head, tail     *DLinkNode
}

type DLinkNode struct {
	key, val  int
	pre, next *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key: key,
		val: value,
	}
}

func Constructor(cap int) LRUCache {
	lru := LRUCache{
		capacity: cap,
		cache:    make(map[int]*DLinkNode, cap),
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (l *LRUCache) Get(key int) int {
	if node, exist := l.cache[key]; exist {
		l.moveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if node, exist := l.cache[key]; exist {
		node.val = value
		l.moveToHead(node)
		return
	}
	newNode := initDLinkNode(key, value)
	l.cache[key] = newNode
	l.addToHead(newNode)
	l.size++
	if l.size > l.capacity {
		tailNode := l.removeTail()
		delete(l.cache, tailNode.key)
		l.size--
	}
}

func (l *LRUCache) moveToHead(node *DLinkNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) addToHead(node *DLinkNode) {
	node.pre = l.head
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
}

func (l *LRUCache) removeTail() *DLinkNode {
	tailNode := l.tail.pre
	l.removeNode(tailNode)
	return tailNode
}

package lru

type LRUCache struct {
	size     int
	capacity int
	key2node map[int]*DNode
	cache    DoubleLinkedList
}

func (l *LRUCache) makeRecently(key int) {
	x := l.key2node[key]
	remove(x)
	l.cache.addLast(x)
}

func (l *LRUCache) addRecently(key, val int) {
	x := &DNode{key, val, nil, nil}
	l.cache.addLast(x)
	l.key2node[key] = x
	l.size++
}

func (l *LRUCache) deleteKey(key int) {
	x := l.key2node[key]
	remove(x)
	delete(l.key2node, key)
	l.size--
}

func (l *LRUCache) removeLeastRecently() {
	deletedNode := l.cache.removeFirst()
	deletedKey := deletedNode.key
	delete(l.key2node, deletedKey)
	l.size--
}

func New(capacity int) *LRUCache {
	head := &DNode{0, 0, nil, nil}
	tail := &DNode{0, 0, nil, nil}
	head.next = tail
	tail.prev = head
	cache := DoubleLinkedList{head, tail}
	key2node := map[int]*DNode{}
	return &LRUCache{0, capacity, key2node, cache}
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.key2node[key]; !ok {
		return -1
	}
	l.makeRecently(key)
	return l.key2node[key].val
}

func (l *LRUCache) Put(key, val int) {
	if _, ok := l.key2node[key]; ok {
		l.deleteKey(key)
		l.addRecently(key, val)
		return
	}

	if l.capacity == l.size {
		l.removeLeastRecently()
	}
	l.addRecently(key, val)
}

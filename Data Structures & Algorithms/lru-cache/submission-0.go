type Node struct{
	key, val int
	prev, next *Node
}

type LRUCache struct {
	capacity int
   cache map[int]*Node
   left, right *Node
   // left is least used
   // right is most used
}

func (this *LRUCache) remove(node *Node){
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *Node){
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.prev = prev
	node.next = next
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		left: &Node{},
		right: &Node{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (this *LRUCache) Get(key int) int {
	// get from the map
	// set it as most used key
	if node, exists := this.cache[key]; exists{
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1  
}

func (this *LRUCache) Put(key int, value int) {
	// exists update, most recent
	if node, exists := this.cache[key]; exists{
		this.remove(node)
		delete(this.cache, key)
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.insert(node)

	// capacity reached, delete lru
	if len(this.cache) > this.capacity{
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
    
}

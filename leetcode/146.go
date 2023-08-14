package leetcode

// https://leetcode.cn/problems/lru-cache/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 146. LRU 缓存
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

// 示例：

// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]

// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4

// 提示：

// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 105
// 最多调用 2 * 105 次 get 和 put

type LRUCache struct {
	cap  int
	size int
	data map[int]*Node
	// queue []int
	Head, Tail *Node
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Pre  *Node
}

func ConstructorLRU(capacity int) LRUCache {
	l := LRUCache{
		cap:  capacity,
		size: 0,
		data: make(map[int]*Node, capacity),
		Head: &Node{Key: -1, Val: -1},
		Tail: &Node{Key: -1, Val: -1},
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head

	return l
}

func (this *LRUCache) Get(key int) int {
	ans := -1
	if node, ok := this.data[key]; ok {
		ans = node.Val
		// 更新至队前
		this.move2Head(node)
	}

	return ans

}

func (this *LRUCache) Put(key int, value int) {
	// 判断是否在里面,在则更新
	if node, ok := this.data[key]; ok {
		node.Val = value
		// 更新至队前
		this.move2Head(node)
		return
	}

	// 不在则新增
	this.size++
	if this.size > this.cap {
		// 先删除队尾节点后插入
		t := this.Tail.Pre
		delete(this.data, t.Key)
		this.removeNode(this.Tail.Pre)
		this.size--
	}

	// 插入元素
	node := &Node{Key: key, Val: value}
	this.data[key] = node
	// this.move2Head(node)
	this.insertNode(node, this.Head)

	return

}

func (this *LRUCache) move2Head(node *Node) {
	this.removeNode(node)
	this.insertNode(node, this.Head)

	return
}

func (this *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	return

}

func (this *LRUCache) insertNode(node, dstNode *Node) {
	// 插入到dstNode之后
	node.Next = dstNode.Next
	dstNode.Next.Pre = node
	node.Pre = dstNode
	dstNode.Next = node

	return

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

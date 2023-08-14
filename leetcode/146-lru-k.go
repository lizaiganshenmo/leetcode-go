package leetcode

// 基于LRU算法扩展实现 LRU-K 算法
// 假定存储数据类型为 key:int ,val:int
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 105

// 采用 LRU-K 模型实现目录的缓存：

// 数据第一次被访问，加入到访问历史列表；

// 如果数据在访问历史列表里后没有达到 K 次访问，则按照一定规则（LRU）淘汰；

// 当访问历史队列中的数据访问次数达到 K 次后，将数据索引从历史队列删除，将数据移到缓存队列中，并缓存此数据，缓存队列重新按照时间排序；

// 缓存数据队列中被再次访问后，重新排序；

// 需要淘汰数据时，淘汰缓存队列中排在末尾的数据，即：淘汰“倒数第 K 次访问离现在最久”的数据。

// 参考文档 : https://juejin.cn/post/6844904049263771662

type LRUKCache struct {
	k           int
	historyList LRUCacheNew
	cache       LRUCacheNew
}

func ConstructorLRUK(k, historyCap, cacheCap int) LRUKCache {
	l := LRUKCache{
		k:           k,
		historyList: ConstructorLRUNew(historyCap),
		cache:       ConstructorLRUNew(cacheCap),
	}

	return l

}

func (l *LRUKCache) get(key int) int {
	// s1 查是否在 cache中
	if tmp, ok := l.cache.Get(key); ok {
		return tmp
	}

	// s2 判断是否在 historyList 中
	if tmp, ok := l.historyList.Get(key); ok {
		// 额外添加判断引用次数是否到k次,到则移至 l.cache
		l.history2Cache()
		return tmp
	}

	return -1
}

func (l *LRUKCache) put(key, val int) {
	// s1 判断cache中是否有
	if _, ok := l.cache.Get(key); ok {
		l.cache.Head.Next.Val = val // l.cache.Get 中有将此节点移动至 head.next   todo :此处可读性不行
		return
	}

	// s2 判断 histortList 中是否有
	if _, ok := l.historyList.Get(key); ok {
		l.historyList.Head.Next.Val = val // l.historyList.Get 中有将此节点移动至 head.next   todo :此处可读性不行
		return
	} else {
		// 没有则新增进入history
		l.historyList.Put(key, val)
	}

	return

}

func (l *LRUKCache) history2Cache() {
	// 引用次数是否到k次,到则移至 l.cache
	node := l.historyList.Head.Next
	if node.Num >= l.k {
		l.historyList.removeNode(node)
		l.cache.insertNode(node, l.cache.Head)
	}
	return

}

/**********************************************************************/

type LRUCacheNew struct {
	cap  int
	size int
	data map[int]*NodeNew
	// queue []int
	Head, Tail *NodeNew
}

type NodeNew struct {
	Key  int
	Val  int
	Num  int // 记录节点数据使用频次
	Next *NodeNew
	Pre  *NodeNew
}

func ConstructorLRUNew(capacity int) LRUCacheNew {
	l := LRUCacheNew{
		cap:  capacity,
		size: 0,
		data: make(map[int]*NodeNew, capacity),
		Head: &NodeNew{Key: -1, Val: -1, Num: 0},
		Tail: &NodeNew{Key: -1, Val: -1, Num: 0},
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head

	return l
}

func (this *LRUCacheNew) Get(key int) (int, bool) {
	if NodeNew, ok := this.data[key]; ok {
		ans := NodeNew.Val
		NodeNew.Num++
		// 更新至队前
		this.move2Head(NodeNew)
		return ans, true
	}

	return -1, false

}

func (this *LRUCacheNew) Put(key int, value int) {
	// 判断是否在里面,在则更新
	if NodeNew, ok := this.data[key]; ok {
		NodeNew.Val = value
		NodeNew.Num++
		// 更新至队前
		this.move2Head(NodeNew)
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
	NodeNew := &NodeNew{Key: key, Val: value, Num: 1}
	this.data[key] = NodeNew
	// this.move2Head(NodeNew)
	this.insertNode(NodeNew, this.Head)

	return

}

func (this *LRUCacheNew) move2Head(NodeNew *NodeNew) {
	this.removeNode(NodeNew)
	this.insertNode(NodeNew, this.Head)

	return
}

func (this *LRUCacheNew) removeNode(NodeNew *NodeNew) {
	NodeNew.Pre.Next = NodeNew.Next
	NodeNew.Next.Pre = NodeNew.Pre

	return

}

func (this *LRUCacheNew) insertNode(NodeNew, dstNodeNew *NodeNew) {
	// 插入到dstNodeNew之后
	NodeNew.Next = dstNodeNew.Next
	dstNodeNew.Next.Pre = NodeNew
	NodeNew.Pre = dstNodeNew
	dstNodeNew.Next = NodeNew

	return

}

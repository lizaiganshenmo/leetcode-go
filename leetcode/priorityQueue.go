package leetcode

import (
	"container/heap"
)

const (
	barrierTag = 1
	hasGone    = 2
)

type Item struct {
	Value    *Loc // 元素的位置
	Index    int  // 元素在堆中的索引位置
	Priority int  // 元素的优先级
}

type Loc struct {
	X, Y int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 根据优先级比较两个元素的大小
	// return pq[i].Priority < pq[j].Priority
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	// 交换两个元素的位置
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	// 向堆中写入元素
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	// 从堆中弹出最高优先级的元素
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // 表示已经弹出
	*pq = old[0 : n-1]
	return item
}

func NewPriorityQueue(size int) PriorityQueue {
	return make(PriorityQueue, 0, size)
}

// 最短路径长度  nums只包含0 -1   -1表示路障
func AStar(nums [][]int, srcX, srcY, dstX, dstY int) int {
	// 1表示节点路障 barrierTag = 1

	costMap := make(map[Loc]int)      // 存储位置节点之前的代价
	locPathMap := make(map[*Loc]*Loc) // 存储位置节点之前的节点
	// 初始化优先队列
	tDis := ManhaDon(srcX, srcY, dstX, dstY)
	priorityQ := NewPriorityQueue(tDis) // 优先级为 起点到 该点的 曼哈顿距离 (小根堆) + 起点到该节点的路径长度
	loc := Loc{X: srcX, Y: srcY}
	heap.Push(&priorityQ, &Item{Value: &loc, Priority: tDis})
	costMap[loc] = 0

	for priorityQ.Len() > 0 {
		node := priorityQ.Pop().(*Item)
		nodeCost := 0
		if val, ok := costMap[*node.Value]; ok {
			nodeCost = val
		}

		// 节点是 目标节点
		if node.Value.X == dstX && node.Value.Y == dstY {
			// 更改nums,便于查看结果是否正确
			t := node.Value
			pre, ok := locPathMap[t]
			for ok {
				nums[pre.X][pre.Y] = hasGone
				pre, ok = locPathMap[pre]
			}

			return nodeCost
		}
		// 不是目标节点
		locs := getNearLocs(nums, node.Value.X, node.Value.Y)
		for _, loc := range locs {
			tmpLoc := loc
			x := loc.X
			y := loc.Y
			if nums[x][y] == barrierTag {
				// 路障不能走
				continue
			}
			newCost := nodeCost + 1
			if val, ok := costMap[tmpLoc]; !ok || newCost < val {
				// 之前没走过 或者  之前走该节点长度 大于当前方式走的路径长度, 则更新节点
				costMap[tmpLoc] = newCost
				manDis := ManhaDon(dstX, dstY, x, y)
				// fmt.Printf("loc is : (%d, %d)  val : %d  newCost :   %d  priority: %d\n", x, y, val, newCost, manDis+newCost)
				heap.Push(&priorityQ, &Item{Value: &tmpLoc, Priority: manDis + newCost})
				locPathMap[&tmpLoc] = node.Value
			}

		}

	}

	return -1

}

// 曼哈顿距离
func ManhaDon(x1, y1, x2, y2 int) int {
	xDis := x1 - x2
	if xDis < 0 {
		xDis = 0 - xDis
	}

	yDis := y1 - y2
	if yDis < 0 {
		yDis = 0 - yDis
	}

	return xDis + yDis
}

// 获取四周节点
func getNearLocs(nums [][]int, x, y int) (ans []Loc) {
	m := len(nums)
	n := len(nums[0])
	// 上
	if x-1 >= 0 {
		ans = append(ans, Loc{X: x - 1, Y: y})
	}
	// 下
	if x+1 < m {
		ans = append(ans, Loc{X: x + 1, Y: y})
	}
	// 左
	if y-1 >= 0 {
		ans = append(ans, Loc{X: x, Y: y - 1})
	}
	// 右
	if y+1 < n {
		ans = append(ans, Loc{X: x, Y: y + 1})
	}

	return
}

type queue struct {
	Item []*Loc
}

func (q *queue) Pop() *Loc {
	if len(q.Item) == 0 {
		return nil
	}
	t := q.Item[0]
	q.Item = q.Item[1:]
	return t
}

func (q *queue) Push(i *Loc) {
	q.Item = append(q.Item, i)
}

func (q *queue) Len() int {
	return len(q.Item)
}

func NewQueue(size int) queue {
	q := queue{
		Item: make([]*Loc, 0, size),
	}
	return q
}

func DfsFind(nums [][]int, srcX, srcY, dstX, dstY int) int {
	var ans int
	if srcX == dstX && srcY == dstY {
		return ans
	}

	queue := NewQueue(1)
	queue.Push(&Loc{X: srcX, Y: srcY})

	for queue.Len() > 0 {
		ans++
		nextQueue := NewQueue(1)

		t := queue.Pop()
		for t != nil {
			nums[t.X][t.Y] = hasGone // 标记已经遍历过

			locs := getNearLocs(nums, t.X, t.Y)
			for _, loc := range locs {
				tmp := loc
				if nums[loc.X][loc.Y] == hasGone || nums[loc.X][loc.Y] == barrierTag {
					continue
				}
				if loc.X == dstX && loc.Y == dstY {
					return ans
				}

				nextQueue.Push(&tmp)

			}
			t = queue.Pop()

		}
		queue = nextQueue

	}

	return -1

}

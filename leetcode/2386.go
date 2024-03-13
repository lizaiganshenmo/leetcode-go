package leetcode

import (
	"container/heap"
)

// https://leetcode.cn/problems/find-the-k-sum-of-an-array/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint
// // 2386. 找出数组的第 K 大和
// 提示
// 困难
// 给你一个整数数组 nums 和一个 正 整数 k 。你可以选择数组的任一 子序列 并且对其全部元素求和。

// 数组的 第 k 大和 定义为：可以获得的第 k 个 最大 子序列和（子序列和允许出现重复）

// 返回数组的 第 k 大和 。

// 子序列是一个可以由其他数组删除某些或不删除元素排生而来的数组，且派生过程不改变剩余元素的顺序。

// 注意：空子序列的和视作 0 。

// 示例 1：

// 输入：nums = [2,4,-2], k = 5
// 输出：2
// 解释：所有可能获得的子序列和列出如下，按递减顺序排列：
// - 6、4、4、2、2、0、0、-2
// 数组的第 5 大和是 2 。
// 示例 2：

// 输入：nums = [1,-2,3,4,-10,12], k = 16
// 输出：10
// 解释：数组的第 16 大和是 10 。

func KSum1(nums []int, k int) int64 {
	// 思路: dfs获取所有可能的子序列和
	n := len(nums)
	// ziNums := []int{0}
	// todo 此题没有完全过， 34/111 内存爆了  ziNums 应当是个容量为k的优先队列
	pq := make(PriorityQueue1, 0)
	heap.Init(&pq)
	// ziNums := make([]int, 1, 2<<(n-1))
	var dfs func(int) //入参: 下标
	sum := 0
	dfs = func(i int) {
		// 定义终止条件
		if i == n {
			return
		}
		// 选择 nums[i]
		sum += nums[i]
		// fmt.Printf("sum is : %d\n", sum)
		t := &Item1{
			Value:    int64(sum),
			Priority: int64(sum),
		}
		heap.Push(&pq, t)
		// ziNums = append(ziNums, sum)
		dfs(i + 1)

		// 不选
		sum -= nums[i] //复位
		dfs(i + 1)

	}

	dfs(0)

	// 排序
	// sort.Slice(ziNums, func(i, j int) bool {
	// 	return ziNums[i] > ziNums[j]
	// })
	var ans *Item1
	for k > 0 {
		t := heap.Pop(&pq)
		// if t == nil {
		// 	return -1
		// }

		ans = t.(*Item1)
		k--
	}
	return ans.Value

}

type Item1 struct {
	Value int64
	// Index    int
	Priority int64
}
type PriorityQueue1 []*Item1

func (pq *PriorityQueue1) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue1) Less(i, j int) bool {
	return (*pq)[i].Priority > (*pq)[j].Priority
}

func (pq *PriorityQueue1) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	// pq[i].Index = j
	// pq[j].Index = i
}

func (pq *PriorityQueue1) Push(x any) {
	item := x.(*Item1)
	// item.Index = pq.Len()
	*pq = (append(*pq, item))
}

func (pq *PriorityQueue1) Pop() any {
	n := pq.Len()
	if n == 0 {
		return &Item1{}
	}
	item := (*pq)[n-1]
	(*pq) = (*pq)[0 : n-1]

	return item
}

// type Item1 struct {
// 	value    int64 // 值
// 	priority int64 // 优先级
// }

// // PriorityQueue implements a priority queue based on a min heap.
// type PriorityQueue1 []*Item1

// // Len returns the length of the priority queue.
// func (pq PriorityQueue1) Len() int { return len(pq) }

// // Less compares two items based on their priorities.
// func (pq PriorityQueue1) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }

// // Swap swaps two items in the priority queue.
// func (pq PriorityQueue1) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// // Push adds an item to the priority queue.
// func (pq *PriorityQueue1) Push(x interface{}) {
// 	item := x.(*Item1)
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue1) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

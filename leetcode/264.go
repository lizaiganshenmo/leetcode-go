package leetcode

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/ugly-number-ii/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=MEDIUM

func nthUglyNumber(n int) int {
	// 思路 : 优先队列
	// 丑数 就是质因子只包含 2、3 和 5 的正整数。

	var (
		ZhiziEles = []int{2, 3, 5}
	)

	pq := NewMyPriorityQueue(n)
	pq.Push(1) //1 通常被视为丑数。

	set := make(map[int]struct{}, n) // 表示某个数是否已经入队,用于去重
	set[1] = struct{}{}

	num := 1 // 已经找到的丑数 个数
	for num <= n {
		t := heap.Pop(pq).(int)
		if num == n {
			return t
		}

		for _, val := range ZhiziEles {
			tmp := val * t
			if _, ok := set[tmp]; !ok {
				heap.Push(pq, tmp)
				set[tmp] = struct{}{}
			}
		}
		num++

	}

	return 0

}

type MyPriorityQueue struct {
	sort.IntSlice
}

func (mpq *MyPriorityQueue) Push(x any) {
	t := x.(int)
	mpq.IntSlice = append(mpq.IntSlice, t)
}

func (mpq *MyPriorityQueue) Pop() any {
	t := mpq.IntSlice[len(mpq.IntSlice)-1]
	mpq.IntSlice = mpq.IntSlice[:len(mpq.IntSlice)-1]

	return t
}

func NewMyPriorityQueue(n int) *MyPriorityQueue {
	return &MyPriorityQueue{
		make(sort.IntSlice, 0, n),
	}

}

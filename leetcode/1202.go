package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/smallest-string-with-swaps/description/
func smallestStringWithSwaps(s string, pairs [][]int) string {
	// 思路 : 并查集
	// pairs 中可以任意互换位置的下标, 那么一定是等价的, 利用并查集进行分组, 然后排序
	n := len(s)
	dsu := NewDsu3(n)
	for _, val := range pairs {
		dsu.Union(val[0], val[1])
	}
	fmt.Printf("dsu.item is: %+v\n", dsu.item)

	mapping := make(map[int][]byte, n)
	for i, val := range dsu.item {
		t := dsu.Find(val)
		mapping[t] = append(mapping[t], s[i])
	}
	fmt.Printf("mapping is: %+v\n", mapping)

	for _, val := range mapping {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}

	arr := make([]byte, n)
	for i := range arr {
		t := dsu.Find(i)
		arr[i] = mapping[t][0]
		mapping[t] = mapping[t][1:]
	}
	return string(arr)

}

type Dsu3 struct {
	item []int // idx 元素数据(本题为pointId) , val 为父节点元素数据
}

func NewDsu3(size int) *Dsu3 {
	item := make([]int, size)
	for i := range item {
		item[i] = i
	}
	return &Dsu3{
		item: item,
	}

}

// Find 找元素的根节点
func (d *Dsu3) Find(x int) int {
	if d.item[x] != x {
		// 不是根节点，根据父节点继续找
		d.item[x] = d.Find(d.item[x])
	}

	return d.item[x]

}

// Union 1
func (d *Dsu3) Union(a, b int) {
	// 两个集合合并，只需将两个集合根节点，其中一个指向另一个即可
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA == rootB {
		return
	}
	d.item[rootA] = rootB
}

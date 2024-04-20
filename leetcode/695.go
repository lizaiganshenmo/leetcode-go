package leetcode

import "fmt"

// https://leetcode.cn/problems/max-area-of-island/description/
func maxAreaOfIsland(grid [][]int) int {
	// 思路:并查集
	m, n := len(grid), len(grid[0])
	dsu := NewDsu2(m * n)

	haveLand := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				haveLand = true
				markSameLands2(grid, dsu, i, j)
			}

		}
	}

	if !haveLand {
		return 0

	}
	return dsu.GetMaxSetSize()

}

func markSameLands2(grid [][]int, dsu *Dsu2, x, y int) {
	m, n := len(grid), len(grid[0])
	// 下节点
	basePos := x*n + y
	if x < m-1 && grid[x+1][y] == 1 {
		newPos := (x+1)*n + y
		dsu.Union(basePos, newPos)
	}
	// 右节点
	if y < n-1 && grid[x][y+1] == 1 {
		newPos := x*n + y + 1
		dsu.Union(basePos, newPos)
	}
}

// Dsu2 并查集
type Dsu2 struct {
	Arr   []int
	Count int
}

// NewDsu2 并查集初始化
func NewDsu2(cap int) *Dsu2 {
	arr := make([]int, cap)
	for i := range arr {
		arr[i] = i
	}

	return &Dsu2{
		Arr:   arr,
		Count: cap,
	}

}

// Find 找根节点
func (d *Dsu2) Find(a int) int {
	if d.Arr[a] != a {
		d.Arr[a] = d.Find(d.Arr[a])
	}

	return d.Arr[a]

}

// Union 节点合并
func (d *Dsu2) Union(a, b int) {
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA != rootB {
		d.Arr[rootA] = rootB

	}

	return
}

// GetMaxSetSize 并查集最大集合尺寸
func (d *Dsu2) GetMaxSetSize() int {
	// fmt.Printf("d.arr is:%+v\n", d.Arr)
	mapping := make(map[int]int)
	for i := range d.Arr {
		mapping[d.Find(d.Arr[i])]++
	}
	// fmt.Printf("d.arr is:%+v\n", d.Arr)
	fmt.Printf("d.mapp is:%+v\n", mapping)

	ans := 1
	// var ansNum int
	for _, v := range mapping {
		if v > ans {
			ans = v
		}
	}

	return ans

}

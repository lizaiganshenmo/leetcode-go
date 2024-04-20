package leetcode

// https://leetcode.cn/problems/number-of-islands/description/
const (
	waterLabel = '0' // 标记为0表示水
	landLabel  = '1' // 标记为1表示陆地
)

// Dsu1 并查集: https://blog.csdn.net/qq_45832961/article/details/124953295
type Dsu1 struct {
	count int
	item  []int // idx 元素数据(本题为pointId) , val 为父节点元素数据
}

// Find 找元素的根节点
func (d *Dsu1) Find(x int) int {
	if d.item[x] != x {
		// 不是根节点，根据父节点继续找
		d.item[x] = d.Find(d.item[x])
	}

	return d.item[x]

}

// Union 1
func (d *Dsu1) Union(a, b int) {
	// 两个集合合并，只需将两个集合根节点，其中一个指向另一个即可
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA == rootB {
		return
	}
	d.count--
	d.item[rootA] = rootB

}

// GetCount 1
func (d *Dsu1) GetCount() int {
	return d.count
}

// NewDsu1 1
func NewDsu1(size int) *Dsu1 {
	item := make([]int, size)
	for i := range item {
		item[i] = i
	}

	return &Dsu1{
		item:  item,
		count: size,
	}
}

func numIslands(grid [][]byte) int {
	// 思路: 并查集
	m, n := len(grid), len(grid[0])
	dsu := NewDsu1(m * n)

	zeroNum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == waterLabel {
				zeroNum++
			} else if grid[i][j] == landLabel {
				markSameLands1(grid, dsu, i, j)
			}

		}
	}

	return dsu.GetCount() - zeroNum

}

func markSameLands1(grid [][]byte, dsu *Dsu1, x, y int) {
	m, n := len(grid), len(grid[0])
	// 下节点
	basePos := x*n + y
	if x < m-1 && grid[x+1][y] == landLabel {
		newPos := (x+1)*n + y
		dsu.Union(basePos, newPos)
	}
	// 右节点
	if y < n-1 && grid[x][y+1] == landLabel {
		newPos := x*n + y + 1
		dsu.Union(basePos, newPos)
	}
}

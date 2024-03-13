package leetcode

// https://leetcode.cn/problems/surrounded-regions/description/

func solve(board [][]byte) {
	// todo 暂未通过 41 / 58
	// 思路 : 并查集
	// 提前设置一个isOuts 遍历四条边界
	const (
		invalidTag = 'O'
		validTag   = 'X'
	)
	m, n := len(board), len(board[0])

	// 遍历 并查集union
	d := NewDsu(m * n)

	// isOuts := make([]bool, m*n)
	for j := 0; j < n; j++ {
		if board[0][j] == invalidTag {
			d.IsOuts[j] = true
		}
		if board[m-1][j] == invalidTag {
			idx := (m-1)*n + j
			d.IsOuts[idx] = true
		}

	}

	for i := 0; i < m; i++ {
		if board[i][0] == invalidTag {
			idx := i * n
			d.IsOuts[idx] = true
		}
		if board[i][n-1] == invalidTag {
			idx := i*n + n - 1
			d.IsOuts[idx] = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == invalidTag {
				union0Tag(board, d, i, j)
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == invalidTag && !d.IsOut(i*n+j) {
				board[i][j] = validTag
			}
		}
	}

}

func union0Tag(board [][]byte, dsu *Dsu, x, y int) {
	var (
		dirs  = [][]int{{0, 1}, {1, 0}} // 向右、向下union
		m     = len(board)
		n     = len(board[0])
		pos   = x*n + y
		isOut = dsu.IsOuts[pos]
	)

	for _, v := range dirs {
		newX, newY := x+v[0], y+v[1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n && board[newX][newY] == board[x][y] {
			newPos := newX*n + newY
			dsu.Union(pos, newPos)
			if isOut {
				dsu.IsOuts[newPos] = true
			}

			if dsu.IsOuts[newPos] {
				dsu.IsOuts[pos] = true
			}
		}
	}
}

// Dsu 并查集
type Dsu struct {
	Data   []int
	IsOuts []bool
}

// NewDsu 新的并查集
func NewDsu(size int) *Dsu {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	return &Dsu{
		Data:   data,
		IsOuts: make([]bool, size),
	}

}

// Find 找根节点
func (d *Dsu) Find(x int) int {

	if d.Data[x] != x {
		d.Data[x] = d.Find(d.Data[x])
		if d.IsOuts[x] {
			d.IsOuts[d.Data[x]] = true
		}
	}

	return d.Data[x]
}

// Union 合并
func (d *Dsu) Union(x, y int) {
	rootX, rootY := d.Find(x), d.Find(y)
	if d.IsOuts[x] || d.IsOuts[y] || d.IsOuts[rootX] || d.IsOuts[rootY] {
		d.IsOuts[x] = true
		d.IsOuts[y] = true
		d.IsOuts[rootX] = true
		d.IsOuts[rootY] = true
	}

	if rootX != rootY {
		d.Data[rootX] = d.Data[rootY]
	}
}

// IsOut 是否是边界
func (d *Dsu) IsOut(x int) bool {
	if d.IsOuts[x] {
		return true
	}
	rootX := d.Find(x)
	if d.IsOuts[rootX] {
		d.IsOuts[x] = true
		return true
	}

	return false

}

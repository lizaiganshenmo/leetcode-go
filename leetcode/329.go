package leetcode

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/description/
var (
	dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func longestIncreasingPath(matrix [][]int) (ans int) {
	// 思路: 先判断是否可能是 起始节点(即该节点小于等于他周围节点)
	// 从可能得起始节点开始,dfs 遍历路径
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	// 自己做这题超时 135/139   参考题解，记录路径遍历
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isValidStartPoint(matrix, i, j) {
				ans = max329(ans, dfs(matrix, i, j, memo))
			}

		}
	}

	return

}

func isValidStartPoint(matrix [][]int, x, y int) bool {
	m, n := len(matrix), len(matrix[0])
	for _, v := range dirs {
		newX, newY := x+v[0], y+v[1]
		if newX < m && newY < n && newX >= 0 && newY >= 0 && matrix[newX][newY] < matrix[x][y] {
			return false
		}
	}

	return true

}

func dfs(matrix [][]int, x, y int, memo [][]int) int {
	if memo[x][y] != 0 {
		return memo[x][y]
	}
	m, n := len(matrix), len(matrix[0])
	memo[x][y]++
	for _, v := range dirs {
		newX, newY := x+v[0], y+v[1]
		if newX < m && newY < n && newX >= 0 && newY >= 0 && matrix[x][y] < matrix[newX][newY] {
			memo[x][y] = max329(memo[x][y], dfs(matrix, newX, newY, memo)+1)
		}

	}

	return memo[x][y]

}

func max329(a, b int) int {
	if a > b {
		return a
	}

	return b
}

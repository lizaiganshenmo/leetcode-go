package leetcode

// https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/description/
func findMaxFish(grid [][]int) int {
	// 思路: bfs
	const (
		waterTag   = 0
		hasGoneTag = -1
	)
	var (
		directions = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	)

	m, n := len(grid), len(grid[0])

	var bfs func(i, j int) int
	bfs = func(i, j int) int {
		var res int
		queue := [][]int{{i, j}}

		res += grid[i][j]
		grid[i][j] = hasGoneTag
		for len(queue) > 0 {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, direc := range directions {
				newX, newY := x+direc[0], direc[1]+y
				if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] > 0 {
					queue = append(queue, []int{newX, newY})
					res += grid[newX][newY]
					grid[newX][newY] = hasGoneTag
				}
			}
		}

		return res

	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				ans = max2658(ans, bfs(i, j))
			}
		}
	}

	return ans

}

func max2658(a, b int) int {
	if a > b {
		return a
	}
	return b
}

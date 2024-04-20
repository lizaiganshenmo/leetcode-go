package leetcode

func islandPerimeter(grid [][]int) int {
	const (
		hasGoneTag = -1
	)

	var (
		directions = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	)

	m, n := len(grid), len(grid[0])
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for _, dire := range directions {
					newX, newY := i+dire[0], j+dire[1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 0 {
						ans++
					}
				}
				if i == 0 {
					ans++
				}
				if i == m-1 {
					ans++
				}
				if j == 0 {
					ans++
				}
				if j == n-1 {
					ans++

				}

			}
		}
	}

	return ans

}

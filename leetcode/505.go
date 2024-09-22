package leetcode

// https://leetcode.cn/problems/the-maze-ii/description/

func shortestDistance(maze [][]int, start []int, destination []int) int {
	// 思路 : bfs
	var (
		directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	)

	m, n := len(maze), len(maze[0])

	stepCount := make([][]int, m) // stepCnt[i][j] 初始位置到 [i,j]的最短步数
	for i := 0; i < m; i++ {
		stepCount[i] = make([]int, n)
		for j := 0; j < n; j++ {
			stepCount[i][j] = -1
		}
	}
	stepCount[start[0]][start[1]] = 0

	queue := [][]int{start}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		for _, val := range directions {
			newX := t[0] + val[0]
			newY := t[1] + val[1]

			step := 0
			// 本题特性: 一条道走到黑
			for newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == 0 {
				newX += val[0]
				newY += val[1]
				step++
			}
			newX -= val[0]
			newY -= val[1]
			if stepCount[newX][newY] != -1 && stepCount[t[0]][t[1]]+step >= stepCount[newX][newY] {
				continue
			}
			queue = append(queue, []int{newX, newY})
			stepCount[newX][newY] = stepCount[t[0]][t[1]] + step
		}

	}

	return stepCount[destination[0]][destination[1]]

}

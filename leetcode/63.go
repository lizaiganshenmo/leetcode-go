package leetcode

// https://leetcode.cn/problems/unique-paths-ii/?envType=study-plan-v2&envId=top-interview-150

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 动态规划
	// dp[i][j] 表示到(i,j)该点 的路径个数
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	const barrierTag = 1 // 1 表示路障
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == barrierTag || obstacleGrid[m-1][n-1] == barrierTag {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一行 第一列最多只能有一种路径
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] != barrierTag && dp[i-1][0] != 0 {
			dp[i][0] = 1
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] != barrierTag && dp[0][j-1] != 0 {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != barrierTag {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0

			}

		}
	}
	// fmt.Printf("dp is : %+v\n", dp)

	return dp[m-1][n-1]
}

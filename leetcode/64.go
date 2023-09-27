package leetcode

import "fmt"

// https://leetcode.cn/problems/minimum-path-sum/?envType=study-plan-v2&envId=top-interview-150

func minPathSum(grid [][]int) int {
	// 思路 : 动态规划
	// dp[i][j] 表示到(i,j)坐标，该点的最小距离
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一列 只能从最上面走过来
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 第一行 只能从左边走过来
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min777(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Printf("dp is : %+v\n", dp)

	return dp[m-1][n-1]

}

func min777(a, b int) int {
	if a < b {
		return a
	}
	return b
}

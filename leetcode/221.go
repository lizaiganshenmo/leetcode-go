package leetcode

// https://leetcode.cn/problems/maximal-square/?envType=study-plan-v2&envId=top-interview-150

func maximalSquare(matrix [][]byte) int {
	// 思路 : 动态规划
	// dp[i][j] 表示以 (i,j)为右下角最大的正方形面积
	// dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) + 1 (matrix[i][j] == 1)
	const validTag = '1'
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var ans int

	// 第一行为1的值都是1
	for i := range matrix[0] {
		if matrix[0][i] == validTag {
			dp[0][i] = 1
			ans = 1
		}
	}
	// 第一列为1的值都是1
	for i := 0; i < m; i++ {
		if matrix[i][0] == validTag {
			dp[i][0] = 1
			ans = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != validTag {
				continue
			}

			dp[i][j] = minn(minn(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			if dp[i][j] > ans {
				ans = dp[i][j]
			}

		}

	}
	return ans * ans

}

func minn(a, b int) int {
	if a < b {
		return a
	}
	return b

}

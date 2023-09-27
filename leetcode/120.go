package leetcode

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/triangle/?envType=study-plan-v2&envId=top-interview-150

func minimumTotal(triangle [][]int) int {
	// 思路 : 动态规划
	// dp[i][j] 表示从顶部到该节点的最短距离
	// dp[i][j] = min(dp[i-1][j-1],dp[i-1][j]) + triangle[i][j]
	const maxNum = math.MaxInt

	m := len(triangle)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = maxNum
		}
	}

	// 第一行等于自身
	dp[0] = triangle[0]
	for i := range triangle {
		if i == 0 {
			continue
		}
		// 第一列只能来自正上层
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min77(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		// 三角形最右边界点只能是左上角元素
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	fmt.Printf("dp is : %+v\n", dp)
	ans := maxNum
	for _, v := range dp[m-1] {
		ans = min77(ans, v)

	}
	return ans

}

func min77(a, b int) int {
	if a < b {
		return a
	}

	return b
}

package leetcode

// https://leetcode.cn/problems/paint-house/?envType=problem-list-v2&envId=dynamic-programming&difficulty=MEDIUM

func minCost(costs [][]int) int {
	// 思路 : 动态规划
	// dp[i][0-2] 表示 第 i 座房子, 选择0(红色) 1(蓝色) 2(绿色) 的最小花费
	// dp[i][0] = min(dp[i-1][1],dp[i-1][2]) + costs[i][0] // 第i座房子涂红色最小花费
	// dp[i][1] = min(dp[i-1][0],dp[i-1][2]) + costs[i][1] // 第i座房子涂蓝色最小花费
	// ...
	const (
		red = iota
		blue
		green
	)

	n := len(costs)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	// dp 边界
	dp[0] = costs[0]

	for i := 1; i < n; i++ {
		dp[i][red] = min256(dp[i-1][blue], dp[i-1][green]) + costs[i][red]
		dp[i][blue] = min256(dp[i-1][red], dp[i-1][green]) + costs[i][blue]
		dp[i][green] = min256(dp[i-1][red], dp[i-1][blue]) + costs[i][green]
	}

	return min256(min256(dp[n-1][red], dp[n-1][blue]), dp[n-1][green])

}

func min256(a, b int) int {
	if a < b {
		return a
	}
	return b
}

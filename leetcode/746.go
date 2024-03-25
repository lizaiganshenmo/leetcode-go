package leetcode

import "fmt"

// https://leetcode.cn/problems/min-cost-climbing-stairs/description/

func minCostClimbingStairs(cost []int) int {
	// 思路： 动态规划
	// dp[i] 表示到达第i级阶梯的最小花费
	// dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min746(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Printf("dp is:%+v\n", dp)

	return dp[n]

}

func min746(a, b int) int {
	if a < b {
		return a
	}

	return b
}

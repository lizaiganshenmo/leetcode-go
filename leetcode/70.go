package leetcode

// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-interview-150

func climbStairs(n int) int {
	// 思路 : 动态规划 dp[i+1] = dp[i] + dp[i-1]
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 2; i < n; i++ {
		dp[i+1] = dp[i] + dp[i-1]
	}
	return dp[n]

}

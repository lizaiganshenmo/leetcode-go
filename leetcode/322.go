package leetcode

import "sort"

// https://leetcode.cn/problems/coin-change/description/?envType=study-plan-v2&envId=top-interview-150

func coinChange(coins []int, amount int) int {
	// 思路 : 动态规划
	// dp[i] 表示筹集到 i 金额最少需要的钱币数目
	// dp[i] = min(dp[i-coins[0]], dp[i-coins[1]], dp[i-coins[2]],...)
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	n := len(coins)
	for _, v := range coins {
		if v <= amount {
			dp[v] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			idx := i - coins[j]
			if idx <= 0 {
				break
			}
			dp[i] = min66(dp[i], dp[idx]+1)
		}

	}

	if dp[amount] == amount+1 {
		return -1

	}
	return dp[amount]

}

func min66(a, b int) int {
	if a < b {
		return a
	}
	return b
}

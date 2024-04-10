package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/coin-change-ii/description/?envType=daily-question&envId=2024-03-25

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		// 当前要选择coins[i]coins[i]coins[i]时，凑出jjj的方案数取决于凑出j−coins[i]j-coins[i]j−coins[i]的方案数，
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func change1(amount int, coins []int) int {
	// 超时了
	// 思路 :dfs
	n := len(coins)

	// path := []int{}
	var ans int
	var dfs func(target, idx int)

	sort.Ints(coins)
	dfs = func(target, idx int) {
		// 递归终止条件
		if n == idx {
			return
		}
		if target == 0 {
			ans++
			return

		}

		if target < coins[idx] {
			return
		}

		// 选当前硬币
		if target >= coins[idx] {
			dfs(target-coins[idx], idx)
		}

		// 不选
		dfs(target, idx+1)

	}

	dfs(amount, 0)

	return ans
}

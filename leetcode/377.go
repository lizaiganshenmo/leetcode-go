package leetcode

import "sort"

// https://leetcode.cn/problems/combination-sum-iv/description/

func combinationSum4(nums []int, target int) int {
	// 思路 : 动态规划
	// dp[i] 表示组合成i 总共有多少种可能
	// dp[i] = sum(dp[i-j]) j->nums... && i-j>=0
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, target+1)
	dp[0] = 1
	sort.Ints(nums)
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if v > i {
				break
			}
			dp[i] += dp[i-v]
		}
	}

	return dp[target]

}

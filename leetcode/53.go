package leetcode

import "math"

// https://leetcode.cn/problems/maximum-subarray/description/

func maxSubArray(nums []int) int {
	// 思路:动态规划
	// dp[i]表示第i个数为结尾时候最大数组和
	// dp[i] = dp[i-1]+nums[i] // 选择第i个数,当nums[i]大于零时
	// dp[i] = nums[i] // 不选择第i个数,当 dp[i]小于零时
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}

	}

	var ans = math.MinInt
	for _, val := range dp {
		if ans < val {
			ans = val
		}
	}

	return ans

}

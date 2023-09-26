package leetcode

import "fmt"

// https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLIS(nums []int) int {
	// 思路 : 动态规划
	// dp[i] 表示 以 i 结尾的子序列的最大长度
	// dp[i] = max (dp[j]) + 1  -> j 区间 [0:i)  && 满足 nums[j] < nums[i]
	var ans int
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max88(dp[i], dp[j]+1)
			}

		}
		ans = max88(ans, dp[i])

	}

	fmt.Printf("dp is : %+v\n", dp)

	return ans

}

func max88(a, b int) int {
	if a > b {
		return a
	}
	return b
}

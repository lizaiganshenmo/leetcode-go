package leetcode

import "sort"

// https://leetcode.cn/problems/delete-and-earn/description/
func deleteAndEarn(nums []int) int {
	// 思路: 动态规划
	// dp[i] 表示取得 值大小为i 时候，最大的利益
	// dp[i] = max(dp[i-1],dp[i-2]+i*cnt(i))
	n := len(nums)
	if n == 0 {
		return 0
	}

	sort.Ints(nums)

	var maxNum int
	cntMp := make(map[int]int) // key->val : elem : cnt
	for _, v := range nums {
		cntMp[v]++
		maxNum = max740(maxNum, v)
	}

	dp := make([]int, maxNum+1)
	dp[1] = cntMp[1]
	for i := 2; i <= maxNum; i++ {
		dp[i] = max740(dp[i-1], dp[i-2]+i*cntMp[i])
	}

	return dp[maxNum]

}

func max740(a, b int) int {
	if a > b {
		return a
	}
	return b
}

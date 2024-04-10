package leetcode

import "math"

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i/description/?envType=daily-question&envId=2024-03-29

func minimumSum(nums []int) int {
	// 思路 : 动态规划
	// dp1[i] 表示 [0,i)区间最小元素值
	// dp2[i] 表示 (i,n-1]区间最小元素
	n := len(nums)
	if n < 3 {
		return 0
	}

	dp1, dp2 := make([]int, n), make([]int, n)
	dp1[1] = nums[0]
	dp2[n-2] = nums[n-1]
	for i := 2; i < n; i++ {
		dp1[i] = min2908(nums[i-1], dp1[i-1])
		dp2[n-1-i] = min2908(nums[n-i], dp2[n-i])

	}

	ans := math.MaxInt
	for i := 1; i < n-1; i++ {
		if nums[i] > dp1[i] && nums[i] > dp2[i] {
			ans = min2908(ans, nums[i]+dp1[i]+dp2[i])
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans

}
func min2908(a, b int) int {
	if a < b {
		return a
	}
	return b
}

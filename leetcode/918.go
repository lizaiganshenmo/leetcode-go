package leetcode

import "math"

// https://leetcode.cn/problems/maximum-sum-circular-subarray/description/

func maxSubarraySumCircular(nums []int) int {
	// 思路 : 动态规划
	// 参考 https://leetcode.cn/problems/maximum-sum-circular-subarray/solutions/2351107/mei-you-si-lu-yi-zhang-tu-miao-dong-pyth-ilqh/
	var (
		maxArrSum      = math.MinInt
		minArrSum      = 0
		sum            = 0
		curMax, curMin = 0, 0
	)

	for _, val := range nums {
		sum += val
		// 以 nums[i-1] 结尾的子数组选或不选（取 max）+ x = 以 x 结尾的最大子数组和
		curMax = max918(curMax, 0) + val
		maxArrSum = max918(curMax, maxArrSum)
		// 以 nums[i-1] 结尾的子数组选或不选（取 min）+ x = 以 x 结尾的最小子数组和
		curMin = min918(curMin, 0) + val
		minArrSum = min918(curMin, minArrSum)

	}
	if sum == minArrSum {
		return maxArrSum
	}

	return max918(maxArrSum, sum-minArrSum)

}

func max918(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min918(a, b int) int {
	if a < b {
		return a
	}
	return b
}

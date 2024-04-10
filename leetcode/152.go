package leetcode

// https://leetcode.cn/problems/maximum-product-subarray/description/
func maxProduct(nums []int) int {
	// 思路 : 动态规划
	// maxDp[i] 表示以 nums[i] 结尾的最大子数组乘积 if nums[i]>=0  maxDp[i] = max(nums[i],maxDp[i-1]*nums[i]) else maxDp[i] = max(nums[i],minDp[i-1]*nums[i])
	// minDp[i] 表示以 nums[i] 结尾的最小子数组乘积 if nums[i]>=0  minDp[i] = min(nums[i],minDp[i-1]*nums[i]) else minDp[i] = min(nums[i],maxDp[i-1]*nums[i])
	n := len(nums)
	if n == 0 {
		return 0
	}

	minDp, maxDp := make([]int, n), make([]int, n)
	minDp[0], maxDp[0] = nums[0], nums[0]
	var ans = maxDp[0]
	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			maxDp[i] = max152(nums[i], maxDp[i-1]*nums[i])
			minDp[i] = min152(nums[i], minDp[i-1]*nums[i])
		} else {
			maxDp[i] = max152(nums[i], minDp[i-1]*nums[i])
			minDp[i] = min152(nums[i], maxDp[i-1]*nums[i])
		}
		ans = max152(ans, maxDp[i])
	}

	return ans

}

func max152(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min152(a, b int) int {
	if a < b {
		return a
	}
	return b
}

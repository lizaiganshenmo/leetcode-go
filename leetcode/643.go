package leetcode

import "math"

// https://leetcode.cn/problems/maximum-average-subarray-i/description/
func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	left, right := 0, 0
	var tmpSum int
	var ans = float64(math.MinInt)
	kf := float64(k)
	for ; right < n; right++ {
		tmpSum += nums[right]
		if right-left+1 == k {
			tmp := float64(tmpSum) / kf
			ans = math.Max(tmp, ans)
			tmpSum -= nums[left]
			left++

		}

	}
	return ans

}

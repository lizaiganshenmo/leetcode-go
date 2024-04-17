package leetcode

// https://leetcode.cn/problems/subarray-product-less-than-k/description/
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	ans := 0

	left, right := 0, 0
	t := 1
	for right < n {
		t *= nums[right]
		for t >= k && left <= right {
			t /= nums[left]
			left++
		}
		ans += right - left + 1 // right - left + 1表示以right下标为子数组右边界的子数组数目
		right++
	}
	return ans

}

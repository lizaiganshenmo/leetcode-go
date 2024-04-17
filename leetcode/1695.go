package leetcode

// https://leetcode.cn/problems/maximum-erasure-value/description/
func maximumUniqueSubarray(nums []int) int {
	// 思路 : 滑动窗口+哈希
	n := len(nums)
	if n == 0 {
		return 0
	}

	mapping := make(map[int]int) // key:val -> 元素值:所在下标
	left, right := 0, 0
	var tmp, ans int
	for ; right < n; right++ {
		tmp += nums[right]
		if idx, ok := mapping[nums[right]]; ok {
			for i := left; i <= idx; i++ {
				tmp -= nums[i]
				delete(mapping, nums[i])
			}
			left = idx + 1
		}
		mapping[nums[right]] = right
		ans = max1695(ans, tmp)
	}
	return ans

}

func max1695(a, b int) int {
	if a > b {
		return a
	}
	return b
}

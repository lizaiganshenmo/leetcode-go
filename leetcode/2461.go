package leetcode

// https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
func maximumSubarraySum(nums []int, k int) int64 {
	// 思路: 滑动窗口+哈希
	n := len(nums)
	if n < k {
		return 0
	}

	left, right := 0, 0
	var tmp, ans int64
	mapping := make(map[int]int, k) // key:val -> 元素值:所在下标
	for ; right < n; right++ {
		tmp += int64(nums[right])
		if idx, ok := mapping[nums[right]]; ok {
			for i := left; i <= idx; i++ {
				tmp -= int64(nums[i])
			}
			if idx+1 > left {
				left = idx + 1
			}

		}
		mapping[nums[right]] = right

		winSize := right - left + 1
		if winSize == k {
			ans = max2461(ans, tmp)
			tmp -= int64(nums[left])
			left++
		}

	}

	return ans

}

func max2461(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

package leetcode

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/description/
func longestSubarray(nums []int) int {
	// 滑动窗口
	var (
		zeroCnt = 0
		// haveZero    = false // 窗口中是否有0
		left, right = 0, 0
	)
	n := len(nums)
	var ans int
	for ; right < n; right++ {
		if nums[right] == 0 {
			zeroCnt++
			for zeroCnt > 1 {
				if nums[left] == 0 {
					zeroCnt--
				}
				left++
			}

		}
		ans = max1493(ans, right-left)

	}

	if zeroCnt == 0 {
		return n - 1
	}
	return ans

}

func max1493(a, b int) int {
	if a > b {
		return a
	}
	return b
}

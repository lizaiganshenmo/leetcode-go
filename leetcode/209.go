package leetcode

// https://leetcode.cn/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150

func minSubArrayLen(target int, nums []int) int {
	// 思路 : 滑动窗口 双指针
	n := len(nums)
	left, right := 0, 1
	ans := n
	sum := nums[0]
	if sum >= target {
		return 1
	}
	for right < n {
		for right < n && sum < target {
			sum += nums[right]
			right++
		}
		if sum >= target {
			ans = min6(ans, right-left+1)
			for sum >= target {
				sum -= nums[left]
				left++
				ans = min6(ans, right-left+1)
			}

		}

	}

	if ans == n+1 {
		ans = 0
	}

	return ans

}

func min6(a, b int) int {
	if a < b {
		return a
	}

	return b
}

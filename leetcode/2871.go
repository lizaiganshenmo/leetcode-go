package leetcode

import "fmt"

// https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/

func maxSubarrays(nums []int) int {
	// 思路: 贪心
	// 遇见可以是 0的 直接划分为一个子数组
	// and := -1 // -1 就是 111...1，和任何数 AND 都等于那个数
	n := len(nums)
	if n == 1 {
		return 1
	}
	right := 0
	cnt := 0
	ans := nums[right]
	preMin := 1
	for right < n {

		ans &= nums[right]
		if ans == 0 {
			cnt++
			fmt.Printf("right: %d\n", right)
			right++
			// left = right
			preMin = 0
			if right < n {
				ans = nums[right]
			}
		} else {
			right++
		}

	}

	if ans >= 1 && preMin == 1 {
		return 1
	}

	return cnt

}

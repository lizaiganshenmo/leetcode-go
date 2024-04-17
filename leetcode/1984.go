package leetcode

import "sort"

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/
func minimumDifference(nums []int, k int) int {
	// 思路: 排序+滑动窗口
	sort.Ints(nums)
	n := len(nums)
	if n < k {
		return 0
	}
	left := 0
	ans := nums[len(nums)-1]
	for ; left <= n-k; left++ {
		tmp := nums[left+k-1] - nums[left]
		if tmp == 0 {
			return 0
		}
		ans = min(ans, tmp)

	}
	return ans

}

func min1984(a, b int) int {
	if a < b {
		return a
	}
	return b
}

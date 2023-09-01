package leetcode

import "strconv"

// https://leetcode.cn/problems/summary-ranges/description/?envType=study-plan-v2&envId=top-interview-150

func summaryRanges(nums []int) []string {
	var ans []string

	n := len(nums)
	left, right := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			right++
		} else {
			ans = append(ans, getRange(nums, left, right))
			left = i
			right = i
		}
	}
	if right < n {
		ans = append(ans, getRange(nums, left, right))
	}

	return ans

}

func getRange(nums []int, left, right int) string {
	if left == right {
		return strconv.Itoa(nums[left])
	}
	return strconv.Itoa(nums[left]) + "->" + strconv.Itoa(nums[right])
}

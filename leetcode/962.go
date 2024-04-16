package leetcode

import "fmt"

// https://leetcode.cn/problems/maximum-width-ramp/description/
func maxWidthRamp(nums []int) int {
	// 思路 : 单调栈
	n := len(nums)
	stack := make([]int, 0, n)
	lefts := make([]int, n)
	for i := 0; i < n; i++ {
		// 单调递减栈
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}

	}
	fmt.Printf("lefts is:%+v\n", lefts)

	var ans int
	for i := n - 1; i > 0; i-- {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = max962(ans, i-j)
		}

	}
	return ans

}

func max962(a, b int) int {
	if a > b {
		return a
	}
	return b
}

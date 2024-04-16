package leetcode

// https://leetcode.cn/problems/maximal-range-that-each-element-is-maximum-in-it/description/
func maximumLengthOfRanges(nums []int) []int {
	// 思路:单调栈
	n := len(nums)
	lefts, rights := make([]int, n), make([]int, n)
	// 两个单调递减栈
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lefts[i] = -1
		} else {
			lefts[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rights[i] = n
		} else {
			rights[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = rights[i] - lefts[i] - 1
	}
	return ans

}

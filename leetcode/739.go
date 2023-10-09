package leetcode

// https://leetcode.cn/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	// 思路 : 单调栈
	// 从右向左遍历, 栈内保持递减
	n := len(temperatures)
	stack := make([]int, 0, n)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		} else {
			ans[i] = 0
		}
		stack = append(stack, i)

	}

	return ans

}

package leetcode

// https://leetcode.cn/problems/maximum-score-of-a-good-subarray/description/

func maximumScore(heights []int, k int) int {
	// 思路 :单调栈
	// 此题与 leetcode 84几乎完全一样 https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
	n := len(heights)
	if n == 0 {
		return 0
	}

	lBounds, rBounds := make([]int, n), make([]int, n)

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lBounds[i] = -1
		} else {
			lBounds[i] = stack[len(stack)-1]
		}
		stack = append(stack, heights[i])
	}

	stack = make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rBounds[i] = n
		} else {
			rBounds[i] = stack[len(stack)-1]
		}
		stack = append(stack, heights[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		if rBounds[i] > k && lBounds[i] < k {
			ans = max1793(ans, (rBounds[i]-lBounds[i]-1)*heights[i])
		}

	}
	return ans

}

func max1793(x, y int) int {
	if x > y {
		return x
	}
	return y
}

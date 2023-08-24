package leetcode

// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-interview-150

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 示例 1：

// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 示例 2：

// 输入：height = [4,2,0,3,2,5]
// 输出：9

// 提示：

// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

func trap(height []int) int {
	// 思路 : 每个下标对应可盛水（宽度为1) =  min(左边前缀最大pre[i], 右边后缀最大(suf[i])) - height[i]
	var ans int
	n := len(height)
	pre := make([]int, n)
	suf := make([]int, n)

	pre[0] = height[0]
	for i := 1; i < n; i++ {
		pre[i] = max3(pre[i-1], height[i])
	}

	suf[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = max3(suf[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		t := min3(pre[i], suf[i]) - height[i]
		if t > 0 {
			ans += t
		}
	}

	return ans

}

func min3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max3(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package leetcode

// https://leetcode.cn/problems/alternating-groups-ii/description/
func numberOfAlternatingGroups(colors []int, k int) int {
	// 思路 : 动态规划
	n := len(colors)

	ans := 0
	cnt := 0 // cnt 表示以i为下标结尾的 交替元素个数, 若cnt > k个，则ans++
	for i := 0; i < n*2; i++ {
		if i > 0 && colors[i%n] == colors[(i-1)%n] {
			cnt = 0
		}
		cnt++

		if i >= n && cnt >= k {
			// i>=n 是因为 数组遍历两次(环形数组缘故), 所以需要保证实际每个数组下标只统计一次
			ans++
		}
	}

	return ans

}

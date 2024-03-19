package leetcode

import "sort"

// https://leetcode.cn/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {
	// 思路 : 贪心算法
	// 首先对数组进行排序
	// 后面遇到重叠区间情况后，优先去除区间更大的那个
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	var ans int
	n := len(intervals)
	idx := 0
	for i := 1; i < n; i++ {
		if intervals[i][0] < intervals[idx][1] {
			ans++
		} else {
			idx = i
		}

	}
	// if idx != n-1 {
	// 	ans++
	// }

	return ans

}

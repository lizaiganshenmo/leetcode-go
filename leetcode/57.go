package leetcode

import "fmt"

// https://leetcode.cn/problems/insert-interval/description/

//区间重叠 开始节点 一定小于等于newInterval的结束值 结束结点一定大于等于newInterval的开始值

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	idx := 0
	n := len(intervals)

	// 左侧未重叠区间添加
	for idx < n && intervals[idx][1] < newInterval[0] {
		ans = append(ans, intervals[idx])
		idx++
	}
	fmt.Printf("ans is: %+v\n", ans)

	// 重叠部分 结束结点一定大于等于newInterval的开始值 && 开始节点 一定小于等于newInterval的结束值
	for idx < n && intervals[idx][1] >= newInterval[0] && intervals[idx][0] <= newInterval[1] {
		fmt.Printf("idx is: %+v\n", idx)
		newInterval[0] = min5(intervals[idx][0], newInterval[0])
		newInterval[1] = max5(intervals[idx][1], newInterval[1])
		idx++
	}
	ans = append(ans, newInterval)
	fmt.Printf("ans is: %+v\n", ans)

	// 右侧未重叠区间添加
	for idx < n && intervals[idx][0] > newInterval[1] {
		ans = append(ans, intervals[idx])
		idx++
	}
	fmt.Printf("ans is: %+v\n", ans)

	return ans

}

func min5(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max5(a, b int) int {
	if a > b {
		return a
	}

	return b
}

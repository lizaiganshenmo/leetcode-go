package leetcode

import (
	"sort"
)

// https://leetcode.cn/contest/weekly-contest-366/problems/minimum-processing-time/

func minProcessingTime(processorTime []int, tasks []int) int {
	// 思路: 贪心 让最早执行的执行任务耗时长的
	sort.Ints(processorTime)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	ans := 0
	n := len(processorTime)
	for i, j := 0, 0; i < n; i++ {
		ans = max10085(ans, processorTime[i]+tasks[j])
		j += 4
	}
	return ans

}

func max10085(a, b int) int {
	if a > b {
		return a
	}
	return b
}

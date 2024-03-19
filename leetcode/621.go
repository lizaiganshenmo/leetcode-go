package leetcode

import "sort"

// https://leetcode.cn/problems/task-scheduler/description/

func leastInterval(tasks []byte, n int) int {
	// 思路 : 贪心
	// 肯定优先排数量最多的在前面,中间出来的空隙则让剩余任务去填充
	taskNums := make([]int, 26)
	for i := range tasks {
		taskNums[tasks[i]-'A']++
	}

	sort.Slice(taskNums, func(i, j int) bool {
		return taskNums[i] > taskNums[j]
	})

	var ans int
	ans = (n+1)*(taskNums[0]-1) + 1
	// extra := ans - taskNums[0]

	for i := 1; i < 26; i++ {
		// extra -= taskNums[i]
		// if extra < 0 {
		// 	ans -= extra
		// 	extra = (n + 1) * (taskNums[i] - 1)
		// }
		if taskNums[i] == taskNums[0] {
			ans++
		}

	}

	return max621(ans, len(tasks))
}

func max621(a, b int) int {
	if a > b {
		return a
	}

	return b
}

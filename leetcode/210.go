package leetcode

// https://leetcode.cn/problems/course-schedule-ii/

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 思路 : bfs
	// 有向无环图, 拓扑排序 从入度为零的点出发遍历
	// 同207题一样
	nodeMap := make(map[int][]int)    // key->val :  先修课程 -> 后学课程列表
	inNums := make([]int, numCourses) // 记录节点的入度数

	for _, v := range prerequisites {
		nodeMap[v[1]] = append(nodeMap[v[1]], v[0])
		inNums[v[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if inNums[i] == 0 {
			q = append(q, i)
		}
	}

	var ans []int
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		ans = append(ans, node)

		for _, v := range nodeMap[node] {
			inNums[v]--
			if inNums[v] == 0 {
				q = append(q, v)
			}
		}

	}

	if len(ans) != numCourses {
		return []int{}
	}

	return ans

}

package leetcode

import "sort"

// https://leetcode.cn/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	// 思路 贪心
	// 尽量分给小的满足要求的
	sort.Ints(g)
	sort.Ints(s)
	var res int
	idx := 0
	for _, v := range g {
		for idx < len(s) && s[idx] < v {
			idx++
		}
		if idx >= len(s) {
			break
		}
		if s[idx] >= v {
			res++
		}
		idx++

	}

	return res

}

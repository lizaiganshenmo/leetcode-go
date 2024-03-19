package leetcode

// https://leetcode.cn/problems/avoid-flood-in-the-city/description/

func avoidFlood(rains []int) []int {
	// todo : 没过 : 47 / 82
	// 思路: 贪心
	// 优先清除最近会再次下雨的湖泊
	const noOperatonTag = -1
	n := len(rains)
	zeroNum := 0
	queue := make([]int, 0)
	mp := make(map[int]struct{})

	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			zeroNum++
			continue
		}

		if _, ok := mp[rains[i]]; ok {
			if zeroNum <= 0 {
				return []int{}
			}
			queue = append(queue, rains[i])
		} else {
			mp[rains[i]] = struct{}{}
		}

	}

	if zeroNum < len(queue) {
		return []int{}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if len(queue) > 0 && rains[i] == 0 {
			ans[i] = queue[0]
			queue = queue[1:]
		} else if rains[i] == 0 {
			ans[i] = 1
		} else {
			ans[i] = noOperatonTag
		}

	}

	return ans

}

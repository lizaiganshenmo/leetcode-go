package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/queue-reconstruction-by-height/description/

func reconstructQueue(people [][]int) [][]int {
	// 参考: https://leetcode.cn/problems/queue-reconstruction-by-height/solutions/486493/xian-pai-xu-zai-cha-dui-dong-hua-yan-shi-suan-fa-g/
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] <= people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}
	return people

}

func insertSlice(slice [][]int, index int, value []int) [][]int {
	if index < 0 || index > len(slice) {
		return slice
	}
	newSlice := make([][]int, len(slice)+1)

	copy(newSlice[:index], slice[:index])
	newSlice[index] = value
	copy(newSlice[index+1:], slice[index:])

	return newSlice
}

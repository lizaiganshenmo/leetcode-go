package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/?envType=study-plan-v2&envId=top-interview-150

func findMinArrowShots(points [][]int) int {
	// 合并区间后的区间数，即为最小引爆数
	// s1 排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
		// if points[i][0] == points[j][0] {
		// 	return points[i][1] < points[j][1]
		// }
		// return points[i][0] < points[j][0]
	})
	fmt.Printf("points is : %+v\n", points)

	right := points[0][1]
	n := len(points)
	ans := 1
	for i := 1; i < n; i++ {
		if right >= points[i][0] {
			continue
		}
		ans++
		right = points[i][1]
	}

	return ans

}

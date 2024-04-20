package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/pond-sizes-lcci/description/

// Pos 位置
type Pos struct {
	X, Y int
}

func pondSizes(land [][]int) []int {
	m, n := len(land), len(land[0])
	// hasGone := make([][]bool, m)
	// for i := range hasGone {
	// 	hasGone[i] = make([]bool, n)
	// }

	var ans []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				ans = append(ans, bfsMark(i, j, land))
			}

		}
	}

	sort.Ints(ans)
	return ans

}

func bfsMark(i, j int, land [][]int) int {
	var (
		directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	)
	m, n := len(land), len(land[0])
	queue := []*Pos{
		{
			X: i,
			Y: j,
		},
	}
	land[i][j] = -1

	var ans int
	for len(queue) > 0 {
		x, y := queue[0].X, queue[0].Y
		// fmt.Printf("queue 0 is:%+v\n", queue[0])
		queue = queue[1:]
		ans++

		for _, dire := range directions {
			newX, newY := x+dire[0], y+dire[1]
			if newX >= 0 && newY >= 0 && newX < m && newY < n && land[newX][newY] == 0 {
				queue = append(queue, &Pos{X: newX, Y: newY})
				land[newX][newY] = -1
			}
		}

	}
	// fmt.Printf("ans is:%d\n", ans)

	return ans

}

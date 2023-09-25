package main

import (
	"daily/leetcode"
	"fmt"
)

func main() {
	nums := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}}

	dst := leetcode.AStar(nums, 0, 0, 6, 10)
	fmt.Printf("dst is : %d\n", dst)
	fmt.Printf("nums is : %+v\n", nums)

	dst1 := leetcode.DfsFind(nums, 0, 0, 6, 10)
	fmt.Printf("dfs dst is : %d\n", dst1)

}

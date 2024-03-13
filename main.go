package main

import "fmt"

func main() {
	// nums := [][]int{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
	// 	{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}}

	// dst := leetcode.AStar(nums, 0, 0, 6, 10)
	// fmt.Printf("dst is : %d\n", dst)
	// fmt.Printf("nums is : %+v\n", nums)

	// dst1 := leetcode.DfsFind(nums, 0, 0, 6, 10)
	// fmt.Printf("dfs dst is : %d\n", dst1)

	var a []int
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Printf("a is : %+v\n", a)

	// var b []int
	b := make([]int, 2)
	copy(b, a)
	fmt.Printf("b : %+v  %d\n", b, len(b))

	original := []int{1, 2, 3}

	// 创建新的空切片
	copied := make([]int, len(original))

	// 将原始切片内容复制到新的切片中
	copy(copied, original[1:2])

	fmt.Println("Original slice:", original) // [1 2 3]
	fmt.Println("Copied slice:", copied)

}

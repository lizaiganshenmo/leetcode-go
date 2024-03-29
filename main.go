package main

import (
	"daily/leetcode"
	"fmt"
)

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

	digits := []int{2, 8, 3}
	fmt.Print("Given digits is: ")
	fmt.Println(digits)

	fmt.Println("n = 2338   expected output : 2333   actual output : ", leetcode.GetMaxNumLessN(digits, 2338))
	fmt.Println("n = 2339   expected output : 2338   actual output : ", leetcode.GetMaxNumLessN(digits, 2339))
	fmt.Println("n = 2399   expected output : 2388   actual output : ", leetcode.GetMaxNumLessN(digits, 2399))
	fmt.Println("n = 2381   expected output : 2338   actual output : ", leetcode.GetMaxNumLessN(digits, 2381))
	fmt.Println("n = 1999   expected output : 888    actual output : ", leetcode.GetMaxNumLessN(digits, 1999))
	fmt.Println("n = 2221   expected output : 888    actual output : ", leetcode.GetMaxNumLessN(digits, 2221))
	fmt.Println("n = 2222   expected output : 888    actual output : ", leetcode.GetMaxNumLessN(digits, 2222))

}

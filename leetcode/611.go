package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/valid-triangle-number/description/

func triangleNumber(nums []int) int {
	// 双指针
	// 三角形条件: a+b>c
	sort.Ints(nums)

	n := len(nums)
	if n < 3 {
		return 0
	}
	// fmt.Printf("nums is:%+v\n", nums)

	var ans int
	for a := 0; a <= n-3; a++ {
		// if a > 0 && nums[a] == nums[a-1] {
		// 	continue
		// }

		for b := a + 1; b <= n-2; b++ {
			// if b > a+1 && nums[b] == nums[b-1] {
			// 	continue
			// }
			t := nums[a] + nums[b]
			for c := b + 1; c <= n-1 && nums[c] < t; c++ {
				// fmt.Printf("trangle is:%d,%d,%d\n", nums[a], nums[b], nums[c])
				ans++
			}

		}
	}

	return ans

}

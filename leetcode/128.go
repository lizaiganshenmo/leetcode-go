package leetcode

import "sort"

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-interview-150

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	// fmt.Printf("n is : %+v\n", n)

	// last := nums[0]
	max, nowMax := 1, 1
	for i := 1; i < n; i++ {
		// fmt.Printf("nums[i] is : %+v\n", nums[i])
		if nums[i] == nums[i-1]+1 {
			// fmt.Printf("111\n")
			nowMax++
			// fmt.Printf("nowMax++, %d\n", nowMax)
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			nowMax = 1
		}
		if max < nowMax {
			max = nowMax
		}

	}
	return max

}

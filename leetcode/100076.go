package leetcode

import (
	"math"
	"sort"
)

// https://leetcode.cn/contest/weekly-contest-365/problems/minimum-size-subarray-in-infinite-array/

// func minSizeSubarray(nums []int, target int) int {
// 	// 思路: 滑动窗口
// 	// 跳出循环 : right == left-1
// 	n := len(nums)
// 	left, right := 0, 0
// 	if nums[left] == target {
// 		return 1
// 	}
// 	for left < n {
// 		if nums[left] < target {
// 			break
// 		}
// 		left++
// 	}

// 	var ans int
// 	length := 1
// 	tmpSum := nums[left]
// 	for tmpSum != target {
// 		for tmpSum < target {
// 			right++
// 			right = right % n
// 			tmpSum += nums[right]
// 			length++

// 		}
// 		if tmpSum == target {
// 			ans = length
// 			break

// 		}
// 		fmt.Printf("len now is : %d\n", length)
// 		for tmpSum > target {
// 			tmpSum -= nums[left]
// 			left++
// 			left = left % n
// 			length--
// 		}
// 		if tmpSum == target {
// 			ans = length
// 			break
// 		}
// 		fmt.Printf("len now is : %d\n", length)

// 	}

// 	return ans

// }

func minSizeSubarray(nums []int, target int) int {
	// 思路: 前缀和
	if nums[0] == target {
		return 1
	}
	sum := 0 // 数组和
	for _, v := range nums {
		sum += v
	}
	extra := 0
	if sum <= target {
		k := target / sum
		if target%sum == 0 {
			return len(nums) * k
		}
		extra = len(nums) * k
		target %= sum
	}
	nums = append(nums, nums...)

	n := len(nums)
	preSums := make([]int, n) // 前缀和
	preSums[0] = nums[0]
	for i := 1; i < n; i++ {
		preSums[i] = preSums[i-1] + nums[i]
	}

	minNum := math.MaxInt
	for i := 0; i < n; i++ {
		if preSums[i] < target {
			continue
		}
		j := sort.Search(n, func(j int) bool {
			return preSums[i]-preSums[j] <= target
		})
		if preSums[i]-preSums[j] == target && i-j < minNum {
			minNum = i - j

		}
		// for j := i - 1; j >= 0; j-- {
		// 	t := preSums[i] - preSums[j]
		// 	if t > target {
		// 		break
		// 	}
		// 	if t == target && i-j < minNum {
		// 		minNum = i - j
		// 	}
		// }
	}

	if minNum == math.MaxInt {
		return -1
	}
	return minNum + extra

}

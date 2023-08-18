package leetcode

import "math/rand"

// https://leetcode.cn/problems/sort-an-array/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 912. 排序数组
// 给你一个整数数组 nums，请你将该数组升序排列。

// 示例 1：

// 输入：nums = [5,2,3,1]
// 输出：[1,2,3,5]
// 示例 2：

// 输入：nums = [5,1,1,2,0,0]
// 输出：[0,0,1,1,2,5]

// 提示：

// 1 <= nums.length <= 5 * 104
// -5 * 104 <= nums[i] <= 5 * 104

func sortArray(nums []int) []int {
	// 快排
	quickSort(nums, 0, len(nums)-1)
	return nums

}

func quickSort(nums []int, left, right int) {
	if right <= left {
		return
	}

	l, r := left, right

	// mid := l + (r-l)/2

	// 随机选择节点
	mid := rand.Intn(right-left) + left
	nums[l], nums[mid] = nums[mid], nums[l]
	for l < r {
		for l < r && nums[r] > nums[left] {
			r--
		}
		for l < r && nums[l] <= nums[left] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]

		}

	}
	nums[l], nums[left] = nums[left], nums[l]
	quickSort(nums, left, l-1)
	quickSort(nums, l+1, right)

}

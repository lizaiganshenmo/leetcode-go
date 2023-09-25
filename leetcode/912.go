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
	// quickSort(nums, 0, len(nums)-1)

	// 堆排
	heapSort(nums)
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

// heapSort
func heapSort(nums []int) {
	// 建堆
	n := len(nums)
	// for i := (n - 1) / 2; i >= 0; i-- {
	// 	heapify(nums, n, i)
	// }

	// // 排序
	// for n > 0 {
	// 	nums[n-1], nums[0] = nums[0], nums[n-1]
	// 	n--
	// 	heapify(nums, n, 0)
	// }

	// 小根堆 自顶向下
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify2(nums, n, i)
	}

	// 排序
	for n > 0 {
		nums[n-1], nums[0] = nums[0], nums[n-1]
		n--
		heapify2(nums, n, 0)
	}

}

func heapify(nums []int, n, idx int) {
	lagest := idx
	left, right := 2*idx+1, 2*idx+2
	if left < n && nums[lagest] < nums[left] {
		lagest = left
	}
	if right < n && nums[lagest] < nums[right] {
		lagest = right
	}

	if lagest != idx {
		nums[lagest], nums[idx] = nums[idx], nums[lagest]
		heapify(nums, n, lagest)
	}
}

func heapify2(nums []int, n, idx int) {
	smallest := idx
	left, right := 2*idx+1, 2*idx+2
	if left < n && nums[smallest] > nums[left] {
		smallest = left
	}
	if right < n && nums[smallest] > nums[right] {
		smallest = right
	}

	if smallest != idx {
		nums[smallest], nums[idx] = nums[idx], nums[smallest]
		heapify(nums, n, smallest)
	}
}

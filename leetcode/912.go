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
	// idx节点的根（父)节点 (idx-1)/2
	n := len(nums)
	for i := (n - 1) / 2; i >= 0; i-- {
		heapfy(nums, i, n)
	}

	// 排序
	for n > 0 {
		nums[n-1], nums[0] = nums[0], nums[n-1]
		n--
		heapfy(nums, 0, n)

	}

}

// 切片 ,当前处理下标, 当前堆大小  大根堆维护
func heapfy(nums []int, idx, n int) {
	lagest := idx
	// idx节点的根（父)节点 (idx-1)/2
	left := 2*idx + 1  //左子节点
	right := 2*idx + 2 // 右子节点

	if left < n && nums[lagest] < nums[left] {
		lagest = left
	}
	if right < n && nums[lagest] < nums[right] {
		lagest = right
	}

	// 当前根节点需要更改，则 更改后的子节点同样需要维护性质
	if lagest != idx {
		nums[lagest], nums[idx] = nums[idx], nums[lagest]
		heapfy(nums, lagest, n)
	}

}

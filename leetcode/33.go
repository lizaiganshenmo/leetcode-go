package leetcode

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=bytedance-2023-fall-sprint

func search(nums []int, target int) int {
	minPos := findMinPos(nums)
	if nums[minPos] > target || (minPos > 0 && nums[minPos-1] < target) {
		return -1
	}
	// 二分找
	var left, right int
	if nums[0] == target {
		return 0

	} else if nums[len(nums)-1] < target {
		left = 0
		right = minPos - 1
	} else {
		left = minPos
		right = len(nums) - 1
	}
	t := right
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == t+1 || nums[left] != target {
		return -1
	}
	return left

}

func findMinPos(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		// 和nums[n-1]比较，如果大  则left = mid+1
		// 小于的话  4 5 6 1 2 3 right
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left

}

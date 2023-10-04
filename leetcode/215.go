package leetcode

// https://leetcode.cn/problems/kth-largest-element-in-an-array/

// 此方法用例过，但超时
// func findKthLargest(nums []int, k int) int {
// 	// 思路 : 分治
// 	// https://leetcode.cn/problems/kth-largest-element-in-an-array/solutions/269218/partition-zhu-shi-by-da-ma-yi-sheng/
// 	left := 0
// 	right := len(nums) - 1

// 	for {
// 		if left >= right { // 重要
// 			return nums[right]
// 		}
// 		p := partition(nums, left, right)
// 		if p+1 == k {
// 			return nums[p]
// 		} else if p+1 < k {
// 			left = p + 1
// 		} else {
// 			right = p - 1
// 		}
// 	}

// }
// func partition(nums []int, left int, right int) int {
// 	l, r := left, right
// 	mid := l + rand.Intn(r-l+1)
// 	tmp := nums[mid]
// 	nums[l], nums[mid] = nums[mid], nums[l]
// 	for l < r {
// 		for l < r && nums[r] < tmp {
// 			r--
// 		}
// 		for l < r && nums[l] >= tmp {
// 			l++
// 		}
// 		if l < r {
// 			nums[l], nums[r] = nums[r], nums[l]
// 		}

// 	}
// 	nums[l], nums[left] = nums[left], nums[l]
// 	return l
// }

// 时间复杂度 O(nlogn)
func findKthLargest(nums []int, k int) int {
	// 思路 : 建堆
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		heapify3(nums, i, n)
	}

	i := 1
	for i <= k {
		nums[0], nums[n-i] = nums[n-i], nums[0]
		heapify3(nums, 0, n-i)
		i++

	}
	return nums[n-k]

}

func heapify3(nums []int, idx, size int) {
	largest := idx
	left, right := 2*idx+1, 2*idx+2
	if left < size && nums[largest] < nums[left] {
		largest = left
	}
	if right < size && nums[largest] < nums[right] {
		largest = right
	}
	if idx != largest {
		nums[idx], nums[largest] = nums[largest], nums[idx]
		heapify3(nums, largest, size)
	}
}

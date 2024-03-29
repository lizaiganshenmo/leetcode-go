package leetcode

// https://leetcode.cn/problems/smallest-k-lcci/description/
func smallestK(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	if k > n {
		//
		return []int{}
	}

	for i := n / 2; i >= 0; i-- {
		heapify4(arr, i, n)
	}

	ans := make([]int, 0, k)
	for k > 0 {
		ans = append(ans, arr[0])
		arr[0], arr[n-1] = arr[n-1], arr[0]
		k--
		n--
		heapify4(arr, 0, n)

	}

	return ans

}

// 小根堆堆化
func heapify4(nums []int, idx, n int) {
	smallest := idx
	lChild, rChild := 2*idx+1, 2*idx+2
	if lChild < n && nums[lChild] < nums[smallest] {
		smallest = lChild
	}
	if rChild < n && nums[rChild] < nums[smallest] {
		smallest = rChild
	}

	if smallest != idx {
		nums[smallest], nums[idx] = nums[idx], nums[smallest]
		heapify4(nums, smallest, n)
	}

}

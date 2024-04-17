package leetcode

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/
func numOfSubarrays(arr []int, k int, threshold int) int {
	// 思路: 滑动窗口
	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] > arr[j]
	// })
	n := len(arr)
	left, right := 0, 0
	var tmpSum int
	var ans int
	for ; right < n; right++ {
		tmpSum += arr[right]
		if right-left+1 == k {
			if tmpSum/k >= threshold {
				ans++
			}
			tmpSum -= arr[left]
			left++

		}

	}

	return ans

}

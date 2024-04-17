package leetcode

// https://leetcode.cn/problems/k-radius-subarray-averages/description/
func getAverages(nums []int, k int) []int {
	// 思路:前缀和
	if k == 0 {
		return nums
	}
	n := len(nums)
	size := 2*k + 1
	// if size > n {
	// 	return []int{}
	// }
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	// fmt.Printf("nums is:%+v\n", preSum)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		l, r := i-k, i+k
		if l < 0 || r >= n {
			ans[i] = -1
		} else {
			ans[i] = (preSum[r+1] - preSum[l]) / size
		}
	}
	return ans

}

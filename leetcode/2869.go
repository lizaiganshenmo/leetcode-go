package leetcode

// https://leetcode.cn/problems/minimum-operations-to-collect-elements/

func minOperations(nums []int, k int) int {
	m := make(map[int]struct{}, k)
	for i := 1; i <= k; i++ {
		m[i] = struct{}{}
	}
	n := len(nums)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if len(m) == 0 {
			break
		}
		if _, ok := m[nums[i]]; ok {
			delete(m, nums[i])
		}
		cnt++
	}

	return cnt

}

package leetcode

// https://leetcode.cn/problems/contains-duplicate-ii/?envType=study-plan-v2&envId=top-interview-150

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int) // val -> idx
	for idx, val := range nums {
		if idx1, ok := mp[val]; ok && idx-idx1 <= k {
			return true
		}
		mp[val] = idx
	}
	return false

}

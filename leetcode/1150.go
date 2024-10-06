package leetcode

// https://leetcode.cn/problems/check-if-a-number-is-majority-element-in-a-sorted-array/description/
func isMajorityElement(nums []int, target int) bool {
	cnt := 0
	for _, val := range nums {
		if val == target {
			cnt++
		}
	}

	return cnt > len(nums)/2

}

package leetcode

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&envId=top-interview-150

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		t := numbers[left] + numbers[right]
		if t == target {
			return []int{left + 1, right + 1}
		} else if t < target {
			left++
		} else {
			right--
		}
	}
	return []int{left + 1, right + 1}

}

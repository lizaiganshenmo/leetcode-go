package leetcode

import "sort"

// https://leetcode.cn/problems/3sum/?envType=study-plan-v2&envId=top-interview-150

func threeSum(nums []int) [][]int {
	n := len(nums)

	var ans [][]int
	if n < 3 {
		return ans
	}

	target := 0
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}
		left, right := i+1, n-1
		for left < right {
			t := nums[left] + nums[right] + nums[i]
			if t == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++

				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if t > target {
				right--
			} else {
				left++
			}

		}
	}

	return ans

}

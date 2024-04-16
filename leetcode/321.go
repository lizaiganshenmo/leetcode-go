package leetcode

import "fmt"

// https://leetcode.cn/problems/create-maximum-number/description/
// todo 没过
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// 思路:单调栈
	// m, n := len(nums1), len(nums2)
	var ans []int
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := getMaxSubsequence(nums1, i)
		s2 := getMaxSubsequence(nums2, k-i)
		fmt.Printf("s1 s2:%+v  %+v\n", s1, s2)
		merged := merge11(s1, s2)
		fmt.Printf("merged:%+v  \n", merged)
		if lexicographicalLess(ans, merged) {
			ans = merged
		}

	}

	return ans

}

func getMaxSubsequence(nums []int, k int) []int {
	delNum := len(nums) - k
	ans := make([]int, 0, k)
	for _, v := range nums {
		for delNum > 0 && len(ans) > 0 && v > ans[len(ans)-1] {
			ans = ans[:len(ans)-1]
			delNum--
		}
		ans = append(ans, v)
	}
	return ans
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge11(a, b []int) []int {
	m, n := len(a), len(b)
	ans := make([]int, 0, m+n)
	idx1, idx2 := 0, 0

	for idx1 < m && idx2 < n {
		if a[idx1] > b[idx2] {
			ans = append(ans, a[idx1])
			idx1++
		} else {
			ans = append(ans, b[idx2])
			idx2++
		}

	}
	for idx1 < m {
		ans = append(ans, a[idx1])
		idx1++
	}
	for idx2 < n {
		ans = append(ans, b[idx2])
		idx2++
	}

	return ans

}

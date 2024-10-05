package leetcode

// https://leetcode.cn/problems/find-common-elements-between-two-arrays/description/
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	for _, val := range nums1 {
		m1[val]++
	}

	for _, val := range nums2 {
		m2[val]++
	}

	var ans1, ans2 int
	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			ans1 += v
		}
	}
	for k, v := range m2 {
		if _, ok := m1[k]; ok {
			ans2 += v
		}
	}

	return []int{ans1, ans2}

}

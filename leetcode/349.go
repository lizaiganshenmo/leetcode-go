package leetcode

// https://leetcode.cn/problems/intersection-of-two-arrays/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 349. 两个数组的交集
// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	m := make(map[int]struct{})
	for i := range nums1 {
		m[nums1[i]] = struct{}{}
	}

	for i := range nums2 {
		if _, ok := m[nums2[i]]; ok {
			ans = append(ans, nums2[i])
			delete(m, nums2[i])
		}
	}

	return ans

}

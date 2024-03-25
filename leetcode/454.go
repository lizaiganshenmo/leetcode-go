package leetcode

// https://leetcode.cn/problems/4sum-ii/description/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// hash表存储 两个数组计算
	mp := make(map[int]int) // sum 个数
	for i := range nums1 {
		for j := range nums2 {
			mp[nums1[i]+nums2[j]]++
		}
	}

	var ans int
	for k := range nums3 {
		for l := range nums4 {
			t := -(nums3[k] + nums4[l])
			ans += mp[t]
		}
	}

	return ans

}

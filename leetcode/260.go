package leetcode

// https://leetcode.cn/problems/single-number-iii/description/

func singleNumber(nums []int) []int {
	// 思路: 哈希
	mp := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			delete(mp, v)
		} else {
			mp[v] = struct{}{}
		}

	}

	ans := make([]int, 0, len(mp))
	for k := range mp {
		ans = append(ans, k)

	}

	return ans

}

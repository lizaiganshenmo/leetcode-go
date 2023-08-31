package leetcode

import "sort"

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-interview-150

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, val := range strs {
		b := []byte(val)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		s := string(b)
		mp[s] = append(mp[s], val)

	}

	var ans [][]string
	for _, val := range mp {
		ans = append(ans, val)
	}

	return ans

}

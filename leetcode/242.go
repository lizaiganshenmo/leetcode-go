package leetcode

// https://leetcode.cn/problems/valid-anagram/?envType=study-plan-v2&envId=top-interview-150

func isAnagram(s string, t string) bool {
	nums := [26]int{}
	for i := range s {
		nums[s[i]-'a']++
	}

	for i := range t {
		nums[t[i]-'a']--
	}

	for _, val := range nums {
		if val != 0 {
			return false
		}
	}

	return true

}

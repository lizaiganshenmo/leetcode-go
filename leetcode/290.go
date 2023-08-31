package leetcode

import "strings"

// https://leetcode.cn/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150

func wordPattern(pattern string, str string) bool {
	str_arr := strings.Fields(str)
	if len(pattern) != len(str_arr) {
		return false
	}
	hash := make(map[byte]string)
	hash2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		v, ok := hash[pattern[i]]
		v2, ok2 := hash2[str_arr[i]]
		if ok && v != str_arr[i] || ok2 && v2 != pattern[i] {
			return false
		} else {
			hash[pattern[i]] = str_arr[i]
			hash2[str_arr[i]] = pattern[i]
		}
	}
	return true
}

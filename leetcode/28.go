package leetcode

import "strings"

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/?envType=study-plan-v2&envId=top-interview-150

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)

}

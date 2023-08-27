package leetcode

import "strings"

//https://leetcode.cn/problems/valid-palindrome/description/?envType=study-plan-v2&envId=top-interview-150

func isPalindrome(s string) bool {
	ss := strings.ToLower(s)
	n := len(ss)
	if n == 1 {
		return true
	}
	for i, j := 0, n-1; i < j; {
		for i < j && !isNumDigit(ss[i]) {
			i++
		}
		for i < j && !isNumDigit(ss[j]) {
			j--
		}
		if i < j {
			if ss[i] != ss[j] {
				return false
			}
			i++
			j--
		}
	}
	return true

}

func isNumDigit(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

package leetcode

// https://leetcode.cn/problems/longest-palindromic-substring/description/?envType=study-plan-v2&envId=top-interview-150

func longestPalindrome(s string) string {
	// 思路 动态规划
	// dp[i][j] 表示字符串 s[i:j+1] 是否是回文串
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	ans := s[0:1]
	maxLen := 1
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for length := 2; length <= n; length++ {
		for start := 0; start < n; start++ {
			end := start + length - 1
			if end >= n {
				break
			}
			if s[start] != s[end] {
				continue
			}
			if length == 2 {
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1]
			}

			if dp[start][end] && length > maxLen {
				maxLen = length
				ans = s[start : end+1]
			}

		}

	}
	return ans

}

func longestPalindrome1(s string) string {
	// 思路 中心扩展法
	var ans string
	var maxLen int

	for i := range s {
		l, r := maxHuiwenAroundCenter(s, i, i)
		tmpLen := r - l + 1
		if tmpLen > maxLen {
			maxLen = tmpLen
			ans = s[l : r+1]
		}
		l, r = maxHuiwenAroundCenter(s, i, i+1)
		tmpLen = r - l + 1
		if tmpLen > maxLen {
			maxLen = tmpLen
			ans = s[l : r+1]
		}

	}

	return ans

}

func maxHuiwenAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

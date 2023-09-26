package leetcode

// https://leetcode.cn/problems/word-break/description/?envType=study-plan-v2&envId=top-interview-150

func wordBreak(s string, wordDict []string) bool {
	// 思路 : 动态规划
	// dp[i] 表示 s[:i] 字符串是否可以被拆分
	// dp[i] = dp[j] && check(s[j:i])
	m := make(map[string]struct{})
	for _, v := range wordDict {
		m[v] = struct{}{}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			_, ok := m[s[j:i]]
			if ok && dp[j] {
				dp[i] = true
				break
			}
		}

	}

	return dp[n]

}

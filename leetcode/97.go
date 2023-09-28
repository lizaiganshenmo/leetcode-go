package leetcode

// https://leetcode.cn/problems/interleaving-string/description/?envType=study-plan-v2&envId=top-interview-150

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 思路  动态规划
	// dp[i][j] 表示 s1[:i] + s2[:j] == s3[:i+j]
	// dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	// 第一行 dp[0][j] 表示是否可以有 s2[:j] == s3[:j]
	for j := 1; j <= n; j++ {
		if s2[:j] != s3[:j] {
			break
		}
		dp[0][j] = true
	}
	// 第一列 dp[i][0] 表示是否可以有 s1[:i] == s3[:i]
	for i := 1; i <= m; i++ {
		if s1[:i] != s3[:i] {
			break
		}
		dp[i][0] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n]

}

package leetcode

// https://leetcode.cn/problems/find-the-substring-with-maximum-cost/description/

func maximumCostSubstring(s string, chars string, vals []int) int {
	// 思路 :动态规划
	// dp[i] 表示字符串 [xx,i)以i结尾的最大开销
	// dp[i] = dp[i-1] + s[i] //s[i]>0
	// dp[i] = dp[i-1]  //s[i]<0
	if len(chars) != len(vals) {
		return 0
	}
	mp := make(map[byte]int, len(chars))
	for i := range chars {
		mp[chars[i]] = vals[i]
	}

	var getKaixiao = func(b byte, mp map[byte]int) int {
		if val, ok := mp[b]; ok {
			return val
		}

		return int(b - 'a' + 1)
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = getKaixiao(s[0], mp)

	ans := dp[0]
	for i := 1; i < n; i++ {

		cnt := getKaixiao(s[i], mp)
		dp[i] = cnt
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] > ans {
			ans = dp[i]
		}

	}

	if ans < 0 {
		return 0
	}

	return ans

}
